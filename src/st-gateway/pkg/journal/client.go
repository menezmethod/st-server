package journal

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"

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
		log.Fatalf("Failed to connect to Journal Service at %s: %v", c.JournalSvcUrl, err)
	}
	log.Printf("Connected to Journal Service at %s", c.JournalSvcUrl)
	return pb.NewJournalServiceClient(cc)
}
