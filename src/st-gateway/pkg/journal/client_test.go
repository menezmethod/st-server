package journal_test

import (
	"context"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"st-gateway/pkg/config"
	"st-gateway/pkg/journal"
	"st-gateway/pkg/journal/pb"
)

var _ = Describe("Client: Journal Service", func() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed loading config", err)
	}

	Describe("Function InitServiceClient()", func() {
		Context("Throwing a random find-one request to the journal service to check connection", func() {
			It("Should not return an error", func() {
				res, err := journal.InitServiceClient(&config.Config{
					Port:          c.Port,
					AuthSvcUrl:    c.AuthSvcUrl,
					JournalSvcUrl: c.JournalSvcUrl,
					ApiVersion:    c.ApiVersion,
				}).FindOneTrade(context.Background(), &pb.FindOneTradeRequest{})
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())
			})
		})
	})
})
