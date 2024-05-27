package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	service := NewUserService(userRepo)

	t.Run("when_delete_user_successful", func(mt *testing.T) {
		id := 999
		userRepo.EXPECT().DeleteUserRepository(id).Return(nil)
		err := service.DeleteUser(id)

		assert.Nil(mt, err)
	})

	t.Run("when_delete_user_error", func(mt *testing.T) {
		id := 999
		userRepo.EXPECT().DeleteUserRepository(id).Return(rest_err.NewBadRequestError("unable create user"))
		err := service.DeleteUser(id)

		assert.NotNil(mt, err)
		assert.Equal(mt, err.Message, "unable create user")
	})
}
