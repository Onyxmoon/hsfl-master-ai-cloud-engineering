package middleware

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth"
	authMock "hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth/_mock"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user"
	userMock "hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user/_mock"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user/model"
	"net/http/httptest"
	"testing"
)

func TestCreateLocalAuthMiddleware(t *testing.T) {
	mockUserRepository := userMock.NewMockRepository(t)
	var userRepository user.Repository = mockUserRepository
	mockTokenGenerator := authMock.NewMockTokenGenerator(t)
	var tokenGenerator auth.TokenGenerator = mockTokenGenerator

	validToken := "valid-token"
	invalidToken := "invalid-token"

	expectedUser := &model.User{
		Id:    1,
		Email: "ada.lovelace@gmail.com",
		Name:  "Ada Lovelace",
		Role:  0,
	}

	middleware := CreateLocalAuthMiddleware(&userRepository, tokenGenerator)

	t.Run("ValidToken", func(t *testing.T) {
		mockUserRepository.EXPECT().FindById(expectedUser.Id).Return(expectedUser, nil).Once()

		claims := map[string]interface{}{"id": float64(expectedUser.Id), "email": expectedUser.Email, "name": expectedUser.Name, "role": float64(expectedUser.Role)}
		mockTokenGenerator.EXPECT().VerifyToken(mock.Anything).Return(claims, nil).Once()

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"GET",
			"/restricted/url", nil)
		request.Header.Set("Authorization", "Bearer "+validToken)

		response := middleware(writer, request)

		assert.NotNil(t, response)
		assert.Equal(t, expectedUser.Id, response.Context().Value("auth_userId"))
		assert.Equal(t, expectedUser.Email, response.Context().Value("auth_userEmail"))
		assert.Equal(t, expectedUser.Name, response.Context().Value("auth_userName"))
		assert.Equal(t, expectedUser.Role, response.Context().Value("auth_userRole"))
	})

	t.Run("InvalidToken", func(t *testing.T) {
		mockTokenGenerator.EXPECT().VerifyToken(mock.Anything).Return(nil, errors.New("invalid token")).Once()
		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"GET",
			"/restricted/url", nil)
		request.Header.Set("Authorization", "Bearer "+invalidToken)

		response := middleware(writer, request)

		assert.NotNil(t, response)
		assert.Nil(t, response.Context().Value("auth_userId"))
		assert.Nil(t, response.Context().Value("auth_userEmail"))
		assert.Nil(t, response.Context().Value("auth_userName"))
		assert.Nil(t, response.Context().Value("auth_userRole"))
	})

	t.Run("InvalidClaims", func(t *testing.T) {
		claims := map[string]interface{}{}
		mockTokenGenerator.EXPECT().VerifyToken(mock.Anything).Return(claims, nil).Once()

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"GET",
			"/restricted/url", nil)
		request.Header.Set("Authorization", "Bearer "+validToken)
		response := middleware(writer, request)

		assert.NotNil(t, response)
		assert.Nil(t, response.Context().Value("auth_userId"))
		assert.Nil(t, response.Context().Value("auth_userEmail"))
		assert.Nil(t, response.Context().Value("auth_userName"))
		assert.Nil(t, response.Context().Value("auth_userRole"))
	})

	t.Run("NonExistentUser", func(t *testing.T) {
		nonExistingUserId := uint64(999)
		claims := map[string]interface{}{"id": float64(nonExistingUserId)}
		mockTokenGenerator.EXPECT().VerifyToken(mock.Anything).Return(claims, nil).Once()
		mockUserRepository.EXPECT().FindById(nonExistingUserId).Return(nil, errors.New(user.ErrorUserNotFound)).Once()

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"GET",
			"/restricted/url", nil)
		request.Header.Set("Authorization", "Bearer "+validToken)
		response := middleware(writer, request)

		assert.NotNil(t, response)
		assert.Nil(t, response.Context().Value("auth_userId"))
		assert.Nil(t, response.Context().Value("auth_userEmail"))
		assert.Nil(t, response.Context().Value("auth_userName"))
		assert.Nil(t, response.Context().Value("auth_userRole"))
	})
}

func Test_getToken(t *testing.T) {
	tests := []struct {
		name          string
		authHeader    string
		expectedToken string
		expectError   bool
	}{
		{"No Authorization Header", "", "", true},
		{"Invalid Format", "InvalidToken", "", true},
		{"Only Bearer", "Bearer", "", true},
		{"Valid Token", "Bearer mytesttoken", "mytesttoken", false},
		{"Extra Parts", "Bearer mytesttoken extra", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "http://example.com", nil)
			if tc.authHeader != "" {
				req.Header.Add("Authorization", tc.authHeader)
			}

			token, err := getToken(req)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error but didn't get one")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if token != tc.expectedToken {
					t.Errorf("Expected token %v, got %v", tc.expectedToken, token)
				}
			}
		})
	}
}
