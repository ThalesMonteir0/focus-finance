package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()
	userRepo := mocks.NewMockUserRepositoryInterface(crtl)
	service := NewUserService(userRepo)

	t.Run("when_user_exists_create_user_return_error", func(mt *testing.T) {
		userDomain := models.NewUserDomain("teste", "teste123@teste123.com", "123456")
		userRepo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUser(userDomain)

		assert.NotNil(mt, err)
		assert.Nil(mt, user)
	})

	t.Run("when_user_not_exists_create_user_return_error", func(mt *testing.T) {
		userDomain := models.NewUserDomain("TESTE", "TESTE1234@TESTE1234.COM", "123456")
		userRepo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		userRepo.EXPECT().CreateUserRepository(userDomain).Return(nil, rest_err.NewBadRequestError("unable create user"))

		user, err := service.CreateUser(userDomain)

		assert.NotNil(mt, err)
		assert.Nil(mt, user)
	})

	t.Run("when_user_not_exists_return_success", func(mt *testing.T) {
		userDomain := models.NewUserDomain("TESTE", "TESTE1234@TESTE1234.COM", "123456")
		userRepo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		userRepo.EXPECT().CreateUserRepository(userDomain).Return(userDomain, nil)

		user, err := service.CreateUser(userDomain)

		assert.Nil(mt, err)
		assert.NotNil(mt, user)
		assert.Equal(mt, userDomain.GetEmail(), user.GetEmail())
		assert.Equal(mt, userDomain.GetName(), user.GetName())
		assert.Equal(mt, userDomain.GetPassword(), user.GetPassword())
	})
}
