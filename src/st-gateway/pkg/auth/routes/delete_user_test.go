package routes_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"st-gateway/pkg/auth/routes"

	"st-gateway/pkg/auth/pb"
	"st-gateway/pkg/auth/pb/mock"
)

var _ = Describe("DeleteUser Route", func() {
	var (
		mockCtrl   *gomock.Controller
		mockClient *mock.MockAuthServiceClient
		router     *gin.Engine
		recorder   *httptest.ResponseRecorder
		id         string
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockClient = mock.NewMockAuthServiceClient(mockCtrl)
		router = gin.Default()
		recorder = httptest.NewRecorder()
		id = "1"

		router.DELETE("/user/:id", func(c *gin.Context) {
			routes.DeleteUser(c, mockClient)
		})
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when the request is valid and the deletion is successful", func() {
		BeforeEach(func() {
			mockClient.EXPECT().
				DeleteUser(gomock.Any(), &pb.DeleteUserRequest{Id: []string{id}}).
				Return(&pb.DeleteUserResponse{Status: 200}, nil)

			req, _ := http.NewRequest("DELETE", "/user/"+id, nil)
			router.ServeHTTP(recorder, req)
		})

		It("should return a status code of 200", func() {
			Expect(recorder.Code).To(Equal(http.StatusOK))
		})

		It("should return a response body with status code 200", func() {
			Expect(recorder.Body.String()).To(Equal(`{"message":"User deleted successfully"}`))
		})
	})

	Context("when the request is invalid", func() {
		Context("when there is no id parameter", func() {
			BeforeEach(func() {
				req, _ := http.NewRequest("DELETE", "/user/", nil)
				router.ServeHTTP(recorder, req)
			})

			It("should return a status code of 404", func() {
				Expect(recorder.Code).To(Equal(http.StatusNotFound))
			})

			It("should return an error message", func() {
				Expect(recorder.Body.String()).To(Equal(`404 page not found`))
			})
		})

		Context("when the server returns an error", func() {
			BeforeEach(func() {
				mockClient.EXPECT().
					DeleteUser(gomock.Any(), &pb.DeleteUserRequest{Id: []string{id}}).
					Return(nil, errors.New("test error"))

				req, _ := http.NewRequest("DELETE", "/user/"+id, nil)
				router.ServeHTTP(recorder, req)
			})

			It("should return a status code of 502", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
