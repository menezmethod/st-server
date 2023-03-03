package routes_test

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	mocks "st-gateway/pkg/auth/pb/mock"
	"st-gateway/pkg/auth/routes"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"st-gateway/pkg/auth/pb"
)

var _ = Describe("Find All Users Route", func() {
	var (
		ctrl *gomock.Controller
		ctx  context.Context
		c    *mocks.MockAuthServiceClient
		rec  *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()
		c = mocks.NewMockAuthServiceClient(ctrl)
		gin.SetMode(gin.TestMode)
		rec = httptest.NewRecorder()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should return a list of users", func() {
		req, err := http.NewRequest("GET", "/users", nil)
		Expect(err).NotTo(HaveOccurred())

		c.EXPECT().FindAllUsers(ctx, &pb.FindAllUsersRequest{}).Return(&pb.FindAllUsersResponse{}, nil)

		router := gin.Default()

		router.GET("/users", func(ctx *gin.Context) {
			routes.FindAllUsers(ctx, c)
		})

		router.ServeHTTP(rec, req)

		Expect(rec.Code).To(Equal(http.StatusOK))
	})
})
