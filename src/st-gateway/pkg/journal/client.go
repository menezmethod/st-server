package journal

import (
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"

	"github.com/menezmethod/st-server/src/st-gateway/configs"
)

type ServiceClient struct {
	JournalClient pb.JournalServiceClient
}

func InitJournalServiceClient(c *configs.Config) pb.JournalServiceClient {
	cc, err := grpc.Dial(c.JournalSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Journal Service at %s: %v", c.JournalSvcUrl, err)
	}
	log.Printf("Connected to Journal Service at %s", c.JournalSvcUrl)
	return pb.NewJournalServiceClient(cc)
}
