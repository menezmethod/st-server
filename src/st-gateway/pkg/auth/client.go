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
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to AuthService: %v", err)
	}

	return pb.NewAuthServiceClient(cc)
}
