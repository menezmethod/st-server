package auth

import (
	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ServiceClient struct {
	AuthServiceClient pb.AuthServiceClient
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
