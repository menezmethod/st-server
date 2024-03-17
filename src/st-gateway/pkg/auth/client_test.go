package auth_test

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"st-gateway/configs"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"st-gateway/pkg/auth"
	"st-gateway/pkg/auth/pb"
)

var _ = Describe("Test InitServiceClient()", func() {

	if err != nil {
		log.Fatalln("failed loading configs", err)
	}
	config := &configs.Config{
		Port:          "8080",
		AuthSvcUrl:    "localhost:50051",
		JournalSvcUrl: "localhost:50052",
		ApiVersion:    "1",
	}

	Context("Register request", func() {

		It("Will send a test register request user and should not return an error", func() {
			res, err := auth.InitServiceClient(config).Register(context.Background(),
				&pb.RegisterRequest{
					Email:     "test@gimenez.com",
					Password:  "123456",
					FirstName: "Test",
					LastName:  "User",
					Bio:       "This user was created by a test",
					Role:      "USER",
					CreatedAt: timestamppb.New(time.Now()),
				})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})

	Context("Login request", func() {
		It("Will log the test user in, get token, and should not return an error", func() {
			res, err := auth.InitServiceClient(config).Login(context.Background(), &pb.LoginRequest{
				Email:    "test@gimenez.com",
				Password: "123456",
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
