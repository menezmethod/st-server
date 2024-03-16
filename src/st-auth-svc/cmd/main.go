package main

import (
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"st-auth-svc/pkg/config"
	"st-auth-svc/pkg/db"
	"st-auth-svc/pkg/pb"
	"st-auth-svc/pkg/services"
	"st-auth-svc/pkg/utils"
)

func main() {
	logger, err := zap.NewDevelopment() // Change to NewProduction in Prod
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Printf("can't sync: %v\n", zap.Error(err))
		}
	}(logger)

	c, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed loading config", zap.Error(err))
	}

	dbHandler := db.InitDB(c.DBUrl, logger)
	if err != nil {
		log.Fatalln("failed to initialize the database", zap.Error(err))
	}

	jwt := initJwt(c)

	reg := prometheus.NewRegistry()
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	reg.MustRegister(grpcMetrics)

	lis, err2 := net.Listen("tcp", c.Port)
	if err2 != nil {
		log.Fatalln("Failed to listen:", err2)
	}

	logger.Info("Auth service listening on", zap.String("port", c.Port))

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)

	grpcMetrics.InitializeMetrics(grpcServer)

	pb.RegisterAuthServiceServer(grpcServer, &services.Server{
		H:      *&dbHandler,
		Logger: logger,
		Jwt:    jwt,
	})

	go func() {
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
		err := http.ListenAndServe(":9093", nil)
		if err != nil {
			log.Printf("can't initialize prometheus metrics port: %v\n", zap.Error(err))
		}
	}()

	if err3 := grpcServer.Serve(lis); err3 != nil {
		log.Fatalln("Failed to serve:", zap.Error(err3))
	}
}

func initJwt(c config.Config) utils.JwtWrapper {
	return utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "st-auth-svc",
		ExpirationHours: 24 * 365,
	}
}
