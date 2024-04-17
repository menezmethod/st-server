package helper

import (
	"context"
	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/helper/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type ServiceClient struct {
	STHelperClient pb.STHelperClient
}

func StHelperClient(c *configs.Config) pb.STHelperClient {
	log.Printf("Initializing gRPC service client for Helper service at URL: %s", c.HelperSvcUrl)

	var cc *grpc.ClientConn
	var err error
	retryInterval := 5 * time.Second
	maxRetryDuration := 60 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), maxRetryDuration)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Exceeded maximum retry duration (%v) for connecting to Helper service", maxRetryDuration)
			return nil
		default:
			cc, err = grpc.DialContext(ctx, c.HelperSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err == nil {
				log.Println("Successfully connected to Helper service")
				return pb.NewSTHelperClient(cc)
			}

			log.Printf("Failed to connect to Helper service at %s: %v. Retrying in %v", c.HelperSvcUrl, err, retryInterval)
			time.Sleep(retryInterval)
			retryInterval = retryInterval + (retryInterval / 5)
		}
	}
}
