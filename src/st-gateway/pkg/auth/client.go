package auth

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"
	"st-gateway/pkg/auth/pb"
	"st-gateway/pkg/config"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	log.Printf("Initializing gRPC service client for Auth service at URL: %s", c.AuthSvcUrl)

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Auth service at %s: %v", c.AuthSvcUrl, err)
	}

	log.Println("Successfully connected to Auth service")

	return pb.NewAuthServiceClient(cc)
}
