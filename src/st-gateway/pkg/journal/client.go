package journal

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"st-gateway/configs"

	"google.golang.org/grpc"
	"st-gateway/pkg/journal/pb"
)

type ServiceClient struct {
	JournalServiceClient pb.JournalServiceClient
	TradeServiceClient   pb.TradeServiceClient
}

func InitJournalServiceClient(c *configs.Config) pb.JournalServiceClient {
	cc, err := grpc.Dial(c.JournalSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Journal Service at %s: %v", c.JournalSvcUrl, err)
	}
	log.Printf("Connected to Journal Service at %s", c.JournalSvcUrl)
	return pb.NewJournalServiceClient(cc)
}

func InitTradeServiceClient(c *configs.Config) pb.TradeServiceClient {
	cc, err := grpc.Dial(c.JournalSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Trade Service at %s: %v", c.JournalSvcUrl, err)
	}
	log.Printf("Connected to Journal Service at %s", c.JournalSvcUrl)
	return pb.NewTradeServiceClient(cc)
}
