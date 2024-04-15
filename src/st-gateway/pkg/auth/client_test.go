package auth_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
)

var _ = Describe("Test AuthServiceClient", func() {
	var (
		config            *configs.Config
		authServiceClient pb.AuthServiceClient
	)

	BeforeEach(func() {
		config = &configs.Config{
			Port:          "8080",
			AuthSvcUrl:    "localhost:50051",
			JournalSvcUrl: "localhost:50052",
			ApiVersion:    "1",
		}

		authServiceClient = auth.InitServiceClient(config)
	})

	Context("Register request", func() {
		It("Will send a test register request user and should not return an error", func() {
			res, err := authServiceClient.Register(context.Background(),
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
			res, err := authServiceClient.Login(context.Background(), &pb.LoginRequest{
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
			res, err := authServiceClient.FindOneUser(context.Background(), &pb.FindOneUserRequest{
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
			res, err := authServiceClient.FindAllUsers(context.Background(), &pb.FindAllUsersRequest{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})
})
