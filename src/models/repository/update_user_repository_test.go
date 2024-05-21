package repository

import (
	"errors"
	"focus-finance/src/configuration/database"
	"focus-finance/src/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_UpdateUserRepository(t *testing.T) {
	db, mock := database.NewMockDB()
	userRepo := NewUserRepository(db)
	expectedSQL := "UPDATE \"users\" SET \"updated_at\"=\\$1,\"name\"=\\$2,\"email\"=\\$3 WHERE \"users\"\\.\"deleted_at\" IS NULL AND \"id\" = \\$4"

	t.Run("updated_user_successful", func(mt *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := userRepo.UpdateUserRepository(models.NewUserUpdateDomain("TESTE", "TESTE@TESTE.COM"), 1)

		assert.Nil(mt, err)
	})

	t.Run("updated_user_error", func(mt *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WithArgs(1).WillReturnError(errors.New("update user failed"))
		mock.ExpectRollback()

		err := userRepo.UpdateUserRepository(models.NewUserUpdateDomain("TESTE", "TESTE@TESTE.COM"), 1)

		assert.NotNil(mt, err)
	})

}
