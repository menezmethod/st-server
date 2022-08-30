package trade

import (
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"st-gateway/pkg/config"
	"st-gateway/pkg/trade/pb"
)

type ServiceClient struct {
	Client pb.TradeServiceClient
}

func InitServiceClient(c *config.Config) pb.TradeServiceClient {
	cc, err := grpc.Dial(c.JournalSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewTradeServiceClient(cc)
}
