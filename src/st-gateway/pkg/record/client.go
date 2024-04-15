package record

import (
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/pb"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"

	"github.com/menezmethod/st-server/src/st-gateway/configs"
)

type ServiceClient struct {
	RecordClient pb.RecordServiceClient
}

func InitRecordServiceClient(c *configs.Config) pb.RecordServiceClient {
	cc, err := grpc.Dial(c.JournalSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Record Service at %s: %v", c.JournalSvcUrl, err)
	}
	log.Printf("Connected to Journal Service at %s", c.JournalSvcUrl)
	return pb.NewRecordServiceClient(cc)
}
