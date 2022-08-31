package auth_test

import (
	"context"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"st-gateway/pkg/auth"
	"st-gateway/pkg/auth/pb"
	"st-gateway/pkg/config"
)

var _ = Describe("Client: Auth Service", func() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed loading config", err)
	}

	Describe("Function InitServiceClient()", func() {
		Context("Throwing a random validate request to the auth service to check connection", func() {
			It("Should not return an error", func() {
				res, err := auth.InitServiceClient(&config.Config{
					Port:          c.Port,
					AuthSvcUrl:    c.AuthSvcUrl,
					JournalSvcUrl: c.JournalSvcUrl,
					ApiVersion:    c.ApiVersion,
				}).Validate(context.Background(), &pb.ValidateRequest{})
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())
			})
		})
	})
})
