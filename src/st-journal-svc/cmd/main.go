package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"st-journal-svc/pkg/config"
	"st-journal-svc/pkg/db"
	"st-journal-svc/pkg/pb"
	"st-journal-svc/pkg/services"
)

func main() {
	logger, err := zap.NewDevelopment() // Change to NewProduction in Prod
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	c, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed loading configs", zap.Error(err))
	}

	h := db.Init(c.DBUrl)
	reg := prometheus.NewRegistry()
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	reg.MustRegister(grpcMetrics)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		logger.Fatal("Failed to listen", zap.String("port", c.Port), zap.Error(err))
	}

	fmt.Println("Journal service listening on:", c.Port)

	serverInstance := &services.Server{
		H:         h,
		Logger:    logger,
		Validator: validator.New(),
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
	)
	pb.RegisterJournalServiceServer(grpcServer, serverInstance)
	pb.RegisterTradeServiceServer(grpcServer, serverInstance)
	grpcMetrics.InitializeMetrics(grpcServer)

	go func() {
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
		err := http.ListenAndServe(":9092", nil)
		if err != nil {
			log.Printf("can't initialize prometheus metrics port: %v\n", zap.Error(err))
		}
	}()

	if err3 := grpcServer.Serve(lis); err3 != nil {
		log.Fatalln("Failed to serve:", zap.Error(err3))
	}
}
