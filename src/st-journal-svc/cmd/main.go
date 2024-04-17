package main

import (
	"fmt"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/auth"
	"log"
	"net"
	"net/http"

	"github.com/go-playground/validator/v10"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/config"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/db"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/services/journal"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/services/record"
)

func main() {
	logger := initLogger()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)

	c := loadConfig(logger)
	h := initDB(c.DBUrl)
	grpcServer := initGRPCServer()

	registerServices(grpcServer, h, logger, c)
	startPrometheusMetrics()

	lis := startListener(c.Port, logger)
	serveGRPC(grpcServer, lis, logger)
}

func initLogger() *zap.Logger {
	logger, err := zap.NewDevelopment() // Change to NewProduction in Prod
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	return logger
}

func loadConfig(logger *zap.Logger) *config.Config {
	c, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed loading configs", zap.Error(err))
	}
	return &c
}

func initDB(dbURL string) db.DB {
	return db.InitDB(dbURL)
}

func initGRPCServer() *grpc.Server {
	reg := prometheus.NewRegistry()
	grpcMetrics := grpcprometheus.NewServerMetrics()
	reg.MustRegister(grpcMetrics)

	return grpc.NewServer(
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
	)
}

func registerServices(grpcServer *grpc.Server, h db.DB, logger *zap.Logger, config *config.Config) {
	authServiceClient := auth.InitAuthServiceClient(config)
	logger.Info("AuthServiceClient initialized successfully")

	pb.RegisterJournalServiceServer(grpcServer, &journal.Server{
		H:                 h,
		Logger:            logger,
		Validator:         validator.New(),
		AuthServiceClient: authServiceClient,
	})
	logger.Info("JournalService registered successfully")

	pb.RegisterRecordServiceServer(grpcServer, &record.Server{
		H:                 h,
		Logger:            logger,
		Validator:         validator.New(),
		AuthServiceClient: authServiceClient,
	})
	logger.Info("RecordService registered successfully")
}

func startPrometheusMetrics() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":9092", nil)
		if err != nil {
			log.Printf("can't initialize prometheus metrics port: %v\n", err)
		}
	}()
}

func startListener(port string, logger *zap.Logger) net.Listener {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("Failed to listen", zap.String("port", port), zap.Error(err))
	}
	fmt.Println("Journal service listening on:", port)
	return lis
}

func serveGRPC(grpcServer *grpc.Server, lis net.Listener, logger *zap.Logger) {
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("Failed to serve", zap.Error(err))
	}
}
