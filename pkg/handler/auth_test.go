package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	trade "github.com/ShatAlex/trading-app"
	"github.com/ShatAlex/trading-app/pkg/service"
	mock_service "github.com/ShatAlex/trading-app/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mock func(s *mock_service.MockAuthorization, user trade.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            trade.User
		mock                 mock
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test", "username":"test", "password":"test"}`,
			inputUser: trade.User{
				Name:     "Test",
				Username: "test",
				Password: "test",
			},
			mock: func(s *mock_service.MockAuthorization, user trade.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Empty field",
			inputBody:            `{"name":"Test", "username":"test"}`,
			mock:                 func(s *mock_service.MockAuthorization, user trade.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"name":"Test", "username":"test", "password":"test"}`,
			inputUser: trade.User{
				Name:     "Test",
				Username: "test",
				Password: "test",
			},
			mock: func(s *mock_service.MockAuthorization, user trade.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mock(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := Handler{services}

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})
	}
}
