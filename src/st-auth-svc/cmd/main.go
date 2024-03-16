package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"st-auth-svc/pkg/config"
	"st-auth-svc/pkg/db"
	"st-auth-svc/pkg/pb"
	"st-auth-svc/pkg/services"
	"st-auth-svc/pkg/utils"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed loading config", err)
	}

	dbHandler := db.InitDB(c.DBUrl)
	if err != nil {
		log.Fatalln("failed to initialize the database", err)
	}

	jwt := initJwt(c)

	startGrpcServer(c.Port, &dbHandler, jwt)
}

func initJwt(c config.Config) utils.JwtWrapper {
	return utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "st-auth-svc",
		ExpirationHours: 24 * 365,
	}
}

func startGrpcServer(port string, dbHandler *db.Handler, jwt utils.JwtWrapper) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	fmt.Println("Auth service listening on:", port)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &services.Server{
		H:   *dbHandler,
		Jwt: jwt,
	})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
