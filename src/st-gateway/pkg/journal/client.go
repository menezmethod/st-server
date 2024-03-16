package journal

import (
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"st-gateway/pkg/config"
	"st-gateway/pkg/journal/pb"
)

type ServiceClient struct {
	Client pb.JournalServiceClient
}

func InitServiceClient(c *config.Config) pb.JournalServiceClient {
	cc, err := grpc.Dial(c.JournalSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewJournalServiceClient(cc)
}
