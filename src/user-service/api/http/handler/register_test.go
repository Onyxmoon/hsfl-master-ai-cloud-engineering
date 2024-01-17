package handler

import (
	"errors"
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

func TestRegisterHandler(t *testing.T) {
	mockUsersRepository := userMock.MockRepository{}
	var userRepository user.Repository = &mockUsersRepository
	mockHasher := cryptoMock.NewMockHasher(t)
	var hasher crypto.Hasher = mockHasher

	registerHandler := NewRegisterHandler(userRepository, hasher)

	t.Run("Valid User (expect 200)", func(t *testing.T) {
		expectedStatus := http.StatusOK
		expectedResponse := `{"id":0,"name":"Grace Hopper","email":"grace.hopper@gmail.com","role":0}`

		userToCreate := &model.User{
			Email:    "grace.hopper@gmail.com",
			Password: []byte("123456"),
			Name:     "Grace Hopper",
			Role:     0,
		}

		mockUsersRepository.EXPECT().FindByEmail("grace.hopper@gmail.com").Return(nil, errors.New(user.ErrorUserNotFound)).Once()
		mockHasher.EXPECT().Hash([]byte("123456")).Return([]byte("123456"), nil).Once()
		mockUsersRepository.EXPECT().Create(userToCreate).Return(userToCreate, nil).Once()

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/register",
			strings.NewReader(`{"email": "grace.hopper@gmail.com", "password": "123456", "name": "Grace Hopper", "role": 0}`),
		)

		registerHandler.Register(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}

		if strings.Compare(expectedResponse, writer.Body.String()) == 0 {
			t.Errorf("Expected response %s, but got %s", expectedResponse, writer.Body.String())
		}

	})

	t.Run("Invalid Request - Empty Password (expect 400)", func(t *testing.T) {
		expectedStatus := http.StatusBadRequest

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/register",
			strings.NewReader(`{"email": "grace.hopper@gmail.com", "name": "Grace Hopper", "role": 0}`),
		)

		registerHandler.Register(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("User already exists (expect 409)", func(t *testing.T) {
		expectedStatus := http.StatusConflict

		userToCreate := &model.User{
			Email:    "grace.hopper@gmail.com",
			Password: []byte("123456"),
			Name:     "Grace Hopper",
			Role:     0,
		}

		mockUsersRepository.EXPECT().FindByEmail("grace.hopper@gmail.com").Return(userToCreate, nil).Once()

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/register",
			strings.NewReader(`{"email": "grace.hopper@gmail.com", "password": "123456", "name": "Grace Hopper", "role": 0}`),
		)

		registerHandler.Register(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("User should not be able to register as admin (expect 403)", func(t *testing.T) {
		expectedStatus := http.StatusForbidden

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/register",
			strings.NewReader(`{"email": "grace.hopper@gmail.com", "password": "123456", "name": "Grace Hopper", "role": 2}`),
		)

		registerHandler.Register(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}

	})

	t.Run("Invalid Request - Empty Password (expect 400)", func(t *testing.T) {
		expectedStatus := http.StatusBadRequest

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/register",
			strings.NewReader(`{"email": "grace.hopper@gmail.com", "password": "", "name": "Grace Hopper", "role": 0}`),
		)

		registerHandler.Register(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("Malformed JSON (expect 400)", func(t *testing.T) {
		expectedStatus := http.StatusBadRequest

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/register",
			strings.NewReader(`{"eil: "grace.hom", ame": "Grace Hopper", "role: 0`),
		)

		registerHandler.Register(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})

	t.Run("Invalid user data, incorrect Type for Email and Password (expected String) (expect 400)", func(t *testing.T) {
		expectedStatus := http.StatusBadRequest

		writer := httptest.NewRecorder()
		request := httptest.NewRequest(
			"POST",
			"/api/v1/authentication/register",
			strings.NewReader(`{"email": 120, "password": 120, "name": "Grace Hopper", "role": 0}`),
		)

		registerHandler.Register(writer, request)

		if writer.Code != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, writer.Code)
		}
	})
}
