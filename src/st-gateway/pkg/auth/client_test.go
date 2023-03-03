package auth_test

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"st-gateway/pkg/auth"
	"st-gateway/pkg/auth/pb"
	"st-gateway/pkg/config"
)

var _ = Describe("Test InitServiceClient()", func() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed loading config", err)
	}
	config := &config.Config{
		Port:          c.Port,
		AuthSvcUrl:    c.AuthSvcUrl,
		JournalSvcUrl: c.JournalSvcUrl,
		ApiVersion:    c.ApiVersion,
	}

	Context("Register request", func() {

		It("Will send a test register request user and should not return an error", func() {
			res, err := auth.InitServiceClient(config).Register(context.Background(),
				&pb.RegisterRequest{
					Email:     wrapperspb.String("test@gimenez.com"),
					Password:  wrapperspb.String("123456"),
					FirstName: wrapperspb.String("Test"),
					LastName:  wrapperspb.String("User"),
					Bio:       wrapperspb.String("This user was created by a test"),
					Role:      wrapperspb.String("USER"),
					CreatedAt: timestamppb.New(time.Now()),
				})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})

	Context("Login request", func() {
		It("Will log the test user in, get token, and should not return an error", func() {
			res, err := auth.InitServiceClient(config).Login(context.Background(), &pb.LoginRequest{
				Email:    wrapperspb.String("test@gimenez.com"),
				Password: wrapperspb.String("123456"),
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
			Expect(res.Data.Token).ToNot(BeEmpty())
		})
	})

	Context("Find Test User request", func() {
		It("Should return test user created about and not return an error", func() {
			res, err := auth.InitServiceClient(config).FindOneUser(context.Background(), &pb.FindOneUserRequest{
				Id: 2,
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
			Expect(res.Data.Email).To(ContainSubstring("test@gimenez.com"))
			Expect(res.Data.FirstName).To(ContainSubstring("Test"))
			Expect(res.Data.LastName).To(ContainSubstring("User"))
			Expect(res.Data.Role).To(ContainSubstring("USER"))
		})
	})

	Context("Find All Users request", func() {
		It("Should not return an error", func() {
			res, err := auth.InitServiceClient(config).FindAllUsers(context.Background(), &pb.FindAllUsersRequest{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})
})
