package routes_test

import (
	"bytes"
	"encoding/json"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/routes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb/mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Register Route", func() {
	var (
		ctrl       *gomock.Controller
		mockClient *mock.MockAuthServiceClient
		r          *gin.Engine
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockClient = mock.NewMockAuthServiceClient(ctrl)

		r = gin.Default()
		r.POST("/register", func(c *gin.Context) {
			routes.Register(c, mockClient)
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Successful registration", func() {
		It("should return status code 201 and a response body with status 'ok'", func() {
			reqBody := routes.RegisterRequestBody{
				Email:     "john.doe@example.com",
				Password:  "password123",
				FirstName: "John",
				LastName:  "Doe",
				Role:      "USER",
			}
			reqBodyBytes, err := json.Marshal(reqBody)
			Expect(err).NotTo(HaveOccurred())

			expectedRes := &pb.RegisterResponse{
				Status: http.StatusOK,
			}
			mockClient.EXPECT().Register(gomock.Any(), &pb.RegisterRequest{
				Email:     reqBody.Email,
				Password:  reqBody.Password,
				FirstName: reqBody.FirstName,
				LastName:  reqBody.LastName,
				Role:      reqBody.Role,
			}).Return(expectedRes, nil)

			req, err := http.NewRequest("POST", "/register", bytes.NewReader(reqBodyBytes))
			Expect(err).NotTo(HaveOccurred())

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))

			var resBody pb.RegisterResponse
			err = json.NewDecoder(w.Body).Decode(&resBody)
			Expect(err).NotTo(HaveOccurred())
			Expect(resBody.Status).To(Equal(expectedRes.Status))
		})
	})
})
