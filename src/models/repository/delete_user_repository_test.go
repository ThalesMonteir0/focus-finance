package repository

import (
	"focus-finance/src/configuration/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_DeleteUserRepository(t *testing.T) {
	db, mock := database.NewMockDB()
	userRepo := NewUserRepository(db)

	expectedSQL := `UPDATE "users" SET "deleted_at"=\$1 WHERE "users"\."id" = \$2 AND "users"\."deleted_at" IS NULL`
	mock.ExpectBegin()
	mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	t.Run("user_deleted_successful", func(mt *testing.T) {
		err := userRepo.DeleteUserRepository(1)

		assert.Nil(mt, err)
	})

}
