package auth

import (
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

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
		fmt.Println("No  connection:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
