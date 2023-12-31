package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/kiryu-dev/mykinolist/internal/model"
	"github.com/kiryu-dev/mykinolist/internal/service"
	mock_service "github.com/kiryu-dev/mykinolist/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestController_IdentifyUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthService, tokens *model.Tokens)
	type testCase struct {
		name                 string
		headerName           string
		headerValue          string
		cookieName           string
		cookieValue          string
		tokens               model.Tokens
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}
	testCases := []testCase{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer sOmEt0kee3N",
			cookieName:  "refreshToken",
			tokens: model.Tokens{
				AccessToken: "sOmEt0kee3N",
			},
			mockBehavior: func(s *mock_service.MockAuthService, tokens *model.Tokens) {
				s.EXPECT().ParseAccessToken(tokens.AccessToken).Return(int64(666), nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "666",
		},
		{
			name:        "Access token out of date",
			headerName:  "Authorization",
			headerValue: "Bearer tokenOutOfDate",
			cookieName:  "refreshToken",
			cookieValue: "ImValidToken",
			tokens: model.Tokens{
				AccessToken:  "tokenOutOfDate",
				RefreshToken: "ImValidToken",
			},
			mockBehavior: func(s *mock_service.MockAuthService, tokens *model.Tokens) {
				var id int64 = 88
				s.EXPECT().ParseAccessToken(tokens.AccessToken).Return(
					id, fmt.Errorf("token expiration date has passed"),
				)
				s.EXPECT().ParseRefreshToken(tokens.RefreshToken).Return(id, nil)
				s.EXPECT().UpdateTokens(id).Return(&model.Tokens{
					AccessToken:  "new_access_token",
					RefreshToken: "new_refresh_token",
				}, nil) // user gets new tokens
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "88",
		},
		{
			name:        "Invalid tokens",
			headerName:  "Authorization",
			headerValue: "Bearer invalidToken",
			cookieName:  "refreshToken",
			tokens: model.Tokens{
				AccessToken: "invalidToken",
			},
			mockBehavior: func(s *mock_service.MockAuthService, tokens *model.Tokens) {
				s.EXPECT().ParseAccessToken(tokens.AccessToken).Return(
					int64(0),
					fmt.Errorf("token is malformed: could not base64 decode header: illegal base64 data at input byte 36"),
				)
			},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: "{\"error\":\"token is malformed: could not base64 decode header: illegal base64 data at input byte 36\"}\n",
		},
		{
			name:        "Access token out of date but there is no refresh token",
			headerName:  "Authorization",
			headerValue: "Bearer tokenOutOfDate",
			cookieName:  "refreshToken",
			tokens: model.Tokens{
				AccessToken: "tokenOutOfDate",
			},
			mockBehavior: func(s *mock_service.MockAuthService, tokens *model.Tokens) {
				var id int64 = 88
				s.EXPECT().ParseAccessToken(tokens.AccessToken).Return(
					id, fmt.Errorf("token expiration date has passed"),
				)
				s.EXPECT().ParseRefreshToken(tokens.RefreshToken).Return(
					int64(0),
					fmt.Errorf("token is malformed: token contains an invalid number of segments"),
				)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "{\"error\":\"token expiration date has passed\"}\n",
		},
		{
			name:                 "No auth header",
			cookieName:           "refreshToken",
			tokens:               model.Tokens{},
			mockBehavior:         func(s *mock_service.MockAuthService, tokens *model.Tokens) {},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: "{\"error\":\"invalid authorization header\"}\n",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAuthService(c)
			tc.mockBehavior(auth, &tc.tokens)
			var (
				service    = &service.Service{AuthService: auth}
				middleware = &authMiddleware{service: service}
				router     = mux.NewRouter()
			)
			router.Use(middleware.identifyUser)
			router.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
				id := r.Context().Value(userIDKey{}).(int64)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(fmt.Sprintf("%d", id)))
			}).Methods(http.MethodGet)
			var (
				w   = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/auth", nil)
			)
			http.SetCookie(w, &http.Cookie{
				Name:     tc.cookieName,
				Value:    tc.cookieValue,
				Path:     "/auth",
				MaxAge:   cookieMaxAge,
				HttpOnly: true,
			})
			req.Header.Add(tc.headerName, tc.headerValue)
			cookieHeader := fmt.Sprintf("%s=%s", w.Result().Cookies()[0].Name, w.Result().Cookies()[0].Value)
			req.Header.Add("Cookie", cookieHeader)
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.Equal(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}
