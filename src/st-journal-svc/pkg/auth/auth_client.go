package auth

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/configs"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
)

type ServiceClient struct {
	AuthServiceClient pb.AuthServiceClient
}

func InitAuthServiceClient(c *configs.Config) pb.AuthServiceClient {
	log.Printf("Initializing gRPC service client for Auth service at URL: %s", c.AuthSvcUrl)

	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Auth service at %s: %v", c.AuthSvcUrl, err)
	}

	log.Println("Successfully connected to Auth service")

	return pb.NewAuthServiceClient(cc)
}
