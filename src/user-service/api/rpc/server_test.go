package rpc

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	proto "hsfl.de/group6/hsfl-master-ai-cloud-engineering/lib/rpc/user"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth"
	authMock "hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth/_mock"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user"
	userMock "hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user/_mock"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user/model"
	"testing"
)

func TestNewUserServiceServer(t *testing.T) {
	mockUsersRepository := userMock.NewMockRepository(t)
	var userRepository user.Repository = mockUsersRepository
	mockTokenGenerator := authMock.NewMockTokenGenerator(t)
	var tokenGenerator auth.TokenGenerator = mockTokenGenerator

	server := NewUserServiceServer(&userRepository, tokenGenerator)

	assert.NotNil(t, server)
	assert.NotNil(t, server.userRepository, userRepository)
	assert.NotNil(t, server.tokenGenerator, tokenGenerator)
}

func TestUserServiceServer_ValidateUserToken(t *testing.T) {
	mockUserRepository := userMock.NewMockRepository(t)
	var userRepository user.Repository = mockUserRepository
	mockTokenGenerator := authMock.NewMockTokenGenerator(t)
	var tokenGenerator auth.TokenGenerator = mockTokenGenerator

	server := NewUserServiceServer(&userRepository, tokenGenerator)

	t.Run("ValidToken", func(t *testing.T) {
		expectedUserId := uint64(1)
		expectedUser := &model.User{Id: expectedUserId, Email: "test@example.com", Name: "Test User", Role: 1}
		mockUserRepository.EXPECT().FindById(expectedUserId).Return(expectedUser, nil).Once()

		claims := map[string]interface{}{"id": float64(expectedUserId)}
		mockTokenGenerator.EXPECT().VerifyToken(mock.Anything).Return(claims, nil).Once()

		response, err := server.ValidateUserToken(context.Background(), &proto.ValidateUserTokenRequest{Token: "valid-token"})

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, expectedUser.Id, response.User.Id)
		assert.Equal(t, expectedUser.Email, response.User.Email)
		assert.Equal(t, expectedUser.Name, response.User.Name)
		assert.Equal(t, int64(expectedUser.Role), response.User.Role)
	})

	t.Run("InvalidToken", func(t *testing.T) {
		mockTokenGenerator.EXPECT().VerifyToken("invalid-token").Return(nil, errors.New("invalid token"))

		_, err := server.ValidateUserToken(context.Background(), &proto.ValidateUserTokenRequest{Token: "invalid-token"})
		assert.Error(t, err)
	})

	t.Run("InvalidClaims", func(t *testing.T) {
		claims := make(map[string]interface{})
		mockTokenGenerator.EXPECT().VerifyToken("token-without-id").Return(claims, nil).Once()

		_, err := server.ValidateUserToken(context.Background(), &proto.ValidateUserTokenRequest{Token: "token-without-id"})
		assert.Error(t, err)
	})

	t.Run("NonExistentUser", func(t *testing.T) {
		nonExistingUserId := uint64(999)
		claims := map[string]interface{}{"id": float64(nonExistingUserId)}
		mockTokenGenerator.EXPECT().VerifyToken("valid-token").Return(claims, nil)
		mockUserRepository.EXPECT().FindById(nonExistingUserId).Return(nil, errors.New("user not found")).Once()

		_, err := server.ValidateUserToken(context.Background(), &proto.ValidateUserTokenRequest{Token: "valid-token"})
		assert.Error(t, err)
	})

}
