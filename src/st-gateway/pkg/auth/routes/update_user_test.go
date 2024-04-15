package routes_test

import (
	"errors"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/routes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb/mock"
)

var _ = Describe("UpdateUser", func() {
	var (
		ctrl       *gomock.Controller
		mockClient *mock.MockAuthServiceClient
		router     *gin.Engine
		testUser   = &pb.User{
			Id:        1,
			Email:     "test@example.com",
			Password:  "testpassword",
			FirstName: "Test",
			LastName:  "User",
			Bio:       "I am a test user",
			Role:      "user",
		}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockClient = mock.NewMockAuthServiceClient(ctrl)

		router = gin.Default()
		router.POST("/users/:id", func(ctx *gin.Context) {
			routes.UpdateUser(ctx, mockClient)
		})

		_ = httptest.NewRecorder()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("should update a user", func() {
		// Set up mock client to return success
		mockClient.EXPECT().UpdateUser(
			gomock.Any(),
			&pb.UpdateUserRequest{
				Id:        testUser.Id,
				Email:     testUser.Email,
				Password:  testUser.Password,
				FirstName: testUser.FirstName,
				LastName:  testUser.LastName,
				Bio:       testUser.Bio,
				Role:      testUser.Role,
			},
		).Return(&pb.UpdateUserResponse{}, nil)

		// Set up request
		requestBody := `{
        "email": "test@example.com",
        "password": "testpassword",
        "firstName": "Test",
        "lastName": "User",
        "bio": "I am a test user",
        "role": "user"
    }`
		request := httptest.NewRequest(
			http.MethodPost,
			"/users/"+strconv.Itoa(int(testUser.Id)),
			strings.NewReader(requestBody),
		)
		request.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder() // Initialize recorder

		// Make request and assert response
		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusCreated))
	})

	It("should return an error if binding JSON fails", func() {
		// Set up request with invalid JSON
		request := httptest.NewRequest(
			http.MethodPost,
			"/users/"+strconv.Itoa(int(testUser.Id)),
			strings.NewReader("{invalid-json}"),
		)
		request.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder() // Initialize recorder

		// Make request and assert response
		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusBadRequest))
	})

	It("should return an error if UpdateUser call fails", func() {
		// Set up mock client to return an error
		mockClient.EXPECT().UpdateUser(
			gomock.Any(),
			&pb.UpdateUserRequest{
				Id:        testUser.Id,
				Email:     testUser.Email,
				Password:  testUser.Password,
				FirstName: testUser.FirstName,
				LastName:  testUser.LastName,
				Bio:       testUser.Bio,
				Role:      testUser.Role,
			},
		).Return(nil, errors.New("update user failed"))

		// Set up request
		requestBody := `{
        "email": "test@example.com",
        "password": "testpassword",
        "firstName": "Test",
        "lastName": "User",
        "bio": "I am a test user",
        "role": "user"
    }`
		request := httptest.NewRequest(
			http.MethodPost,
			"/users/"+strconv.Itoa(int(testUser.Id)),
			strings.NewReader(requestBody),
		)
		request.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// Make request and assert response
		router.ServeHTTP(recorder, request)
		Expect(recorder.Code).To(Equal(http.StatusBadGateway))
	})
})
