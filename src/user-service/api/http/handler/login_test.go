package handler

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth"
	authMock "hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth/_mock"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/crypto"
	cryptoMock "hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/crypto/_mock"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user"
	userMock "hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user/_mock"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	mockUsersRepository := userMock.MockRepository{}
	var userRepository user.Repository = &mockUsersRepository
	mockHasher := cryptoMock.NewMockHasher(t)
	var hasher crypto.Hasher = mockHasher
	mockTokenGenerator := authMock.MockTokenGenerator{}
	var tokenGenerator auth.TokenGenerator = &mockTokenGenerator

	loginHandler := NewLoginHandler(userRepository, hasher, tokenGenerator)

	t.Run("Valid User (expect 200)", func(t *testing.T) {
		expectedStatus := http.StatusOK

		mockUsersRepository.EXPECT().FindByEmail("ada.lovelace@gmail.com").Return(&model.User{
			Id:       1,
			Email:    "ada.lovelace@gmail.com",
			Password: []byte("123456"),
			Name:     "Ada Lovelace",
			Role:     model.Customer,
		}, nil).Once()

		mockTokenGenerator.EXPECT().CreateToken(mock.Anything).Return("valid-token", nil).Once()

		mockHasher.EXPECT().Validate([]byte("123456"), []byte("123456")).Return(true).Once()

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/login",
			strings.NewReader(`{"email": "ada.lovelace@gmail.com", "password": "123456"}`),
		)

		loginHandler.Login(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("Invalid User Mail (expect 500)", func(t *testing.T) {
		expectedStatus := http.StatusInternalServerError

		mockUsersRepository.EXPECT().FindByEmail("ada.lovelace@gmail.com").Return(nil, errors.New(user.ErrorUserNotFound)).Once()

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/login",
			strings.NewReader(`{"email": "ada.lovelace@gmail.com", "password": "123456"}`),
		)

		loginHandler.Login(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("Invalid Request - Empty Password (expect 400)", func(t *testing.T) {
		expectedStatus := http.StatusBadRequest

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/login",
			strings.NewReader(`{"email": "ada.lovelace@gmail.com"}`),
		)

		loginHandler.Login(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("Wrong password (expect 401)", func(t *testing.T) {
		expectedStatus := http.StatusUnauthorized

		mockUsersRepository.EXPECT().FindByEmail("ada.lovelace@gmail.com").Return(&model.User{
			Id:       1,
			Email:    "ada.lovelace@gmail.com",
			Password: []byte("894544"),
			Name:     "Ada Lovelace",
			Role:     model.Customer,
		}, nil)

		mockTokenGenerator.EXPECT().CreateToken(mock.Anything).Return("valid-token", nil).Once()

		mockHasher.EXPECT().Validate([]byte("123456"), []byte("894544")).Return(false).Once()

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/login",
			strings.NewReader(`{"email": "ada.lovelace@gmail.com", "password": "123456"}`),
		)

		loginHandler.Login(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("Malformed JSON (expect 400)", func(t *testing.T) {
		expectedStatus := http.StatusBadRequest

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/login",
			strings.NewReader(`{"email: "ada.lovelace@gmail.com`),
		)

		loginHandler.Login(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("Missing field (expect 400)", func(t *testing.T) {
		expectedStatus := http.StatusBadRequest

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/login",
			strings.NewReader(`{"password": "12345"}`),
		)

		loginHandler.Login(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("Invalid user data, incorrect Type for Email and Password (expected String) (expect 400)", func(t *testing.T) {
		expectedStatus := http.StatusBadRequest

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/login",
			strings.NewReader(`{"email": 120 "password": 120}`),
		)

		loginHandler.Login(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})
}
