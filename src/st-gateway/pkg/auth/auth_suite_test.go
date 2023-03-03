package auth_test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"st-gateway/pkg/auth/pb"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Suite")
}

var (
	_      pb.AuthServiceClient
	_      context.Context
	cancel context.CancelFunc
	conn   *grpc.ClientConn
	err    error
)

var _ = BeforeSuite(func() {
	conn, err = grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	Expect(err).ShouldNot(HaveOccurred())

	_, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	_ = pb.NewAuthServiceClient(conn)
})

var _ = AfterSuite(func() {
	cancel()
	err := conn.Close()
	if err != nil {
		return
	}
})
