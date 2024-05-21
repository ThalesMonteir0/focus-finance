package repository

import (
	"errors"
	"focus-finance/src/configuration/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_DeleteUserRepository(t *testing.T) {
	db, mock := database.NewMockDB()
	userRepo := NewUserRepository(db)
	expectedSQL := `UPDATE "users" SET "deleted_at"=\$1 WHERE "users"\."id" = \$2 AND "users"\."deleted_at" IS NULL`

	t.Run("user_deleted_successful", func(mt *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		err := userRepo.DeleteUserRepository(1)

		assert.Nil(mt, err)
	})

	t.Run("user_deleted_error", func(mt *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WithArgs(1).WillReturnError(errors.New("delete failed"))
		mock.ExpectRollback()

		err := userRepo.DeleteUserRepository(1)

		assert.NotNil(mt, err)
	})

}
