package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"st-journal-svc/pkg/config"
	"st-journal-svc/pkg/db"
	"st-journal-svc/pkg/pb"
	"st-journal-svc/pkg/services"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed loading config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Trade service listening on: ", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterJournalServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
