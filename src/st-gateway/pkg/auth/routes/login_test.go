package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"st-gateway/pkg/auth/routes"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"st-gateway/pkg/auth/pb"
	"st-gateway/pkg/auth/pb/mock"
)

var _ = Describe("Login Route", func() {
	var (
		ctrl       *gomock.Controller
		mockClient *mock.MockAuthServiceClient
		r          *gin.Engine
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockClient = mock.NewMockAuthServiceClient(ctrl)

		r = gin.Default()
		r.POST("/login", func(c *gin.Context) {
			routes.Login(c, mockClient)
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Successful login", func() {
		It("should return status code 201 and a response body with status 'ok'", func() {
			reqBody := routes.LoginRequestBody{
				Email:    "john.doe@example.com",
				Password: "password123",
			}
			reqBodyBytes, err := json.Marshal(reqBody)
			Expect(err).NotTo(HaveOccurred())

			expectedRes := &pb.LoginResponse{
				Status: http.StatusOK,
			}
			mockClient.EXPECT().Login(gomock.Any(), &pb.LoginRequest{
				Email:    reqBody.Email,
				Password: reqBody.Password,
			}).Return(expectedRes, nil)

			req, err := http.NewRequest("POST", "/login", bytes.NewReader(reqBodyBytes))
			Expect(err).NotTo(HaveOccurred())

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))

			var resBody pb.LoginResponse
			err = json.NewDecoder(w.Body).Decode(&resBody)
			Expect(err).NotTo(HaveOccurred())
			Expect(resBody.Status).To(Equal(expectedRes.Status))
		})
	})
})
