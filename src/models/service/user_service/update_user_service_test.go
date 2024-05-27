package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	service := NewUserService(userRepo)

	t.Run("when_update_user_success", func(mt *testing.T) {
		domain := models.NewUserDomain("TESTE", "TESTE@TESTE.COM", "123456")
		id := 999
		userRepo.EXPECT().UpdateUserRepository(domain, id).Return(nil)
		err := service.UpdateUser(domain, id)

		assert.Nil(mt, err)
	})

	t.Run("when_update_user_error", func(mt *testing.T) {
		domain := models.NewUserDomain("TESTE", "TESTE@TESTE.COM", "123456")
		id := 999
		userRepo.EXPECT().UpdateUserRepository(domain, id).Return(rest_err.NewInternalServerError("unable update user"))
		err := service.UpdateUser(domain, id)

		assert.NotNil(mt, err)
		assert.Equal(mt, err.Message, "unable update user")
	})
}
