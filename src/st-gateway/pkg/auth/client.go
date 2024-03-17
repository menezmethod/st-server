package auth

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"st-gateway/configs"

	"google.golang.org/grpc"
	"st-gateway/pkg/auth/pb"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *configs.Config) pb.AuthServiceClient {
	log.Printf("Initializing gRPC service client for Auth service at URL: %s", c.AuthSvcUrl)

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Auth service at %s: %v", c.AuthSvcUrl, err)
	}

	log.Println("Successfully connected to Auth service")

	return pb.NewAuthServiceClient(cc)
}
