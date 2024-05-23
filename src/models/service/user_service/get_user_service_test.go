package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"focus-finance/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserService_GetUserByID(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	userRepo := mocks.NewMockUserRepositoryInterface(controller)
	userService := NewUserService(userRepo)

	t.Run("when_exists_an_user_return_success", func(mt *testing.T) {
		id := 999
		userDomain := models.NewUserDomain("TESTE", "TESTE@GMAIL.COM", "123456")
		userDomain.SetID(id)
		userRepo.EXPECT().FindUserByID(id).Return(userDomain, nil)

		userDomainReturn, err := userService.GetUserByID(id)

		assert.Nil(mt, err)
		assert.NotNil(mt, userDomainReturn)
		assert.Equal(mt, userDomainReturn.GetID(), id)
	})

	t.Run("when_does_not_exists_an_user_return_error", func(mt *testing.T) {
		id := 999
		userRepo.EXPECT().FindUserByID(id).Return(nil, rest_err.NewInternalServerError("unable get user"))

		userDomainReturn, err := userService.GetUserByID(id)

		assert.Nil(mt, userDomainReturn)
		assert.NotNil(mt, err)
	})
}
