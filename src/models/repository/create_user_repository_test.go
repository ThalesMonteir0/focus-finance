package repository

import (
	"focus-finance/src/configuration/database"
	"focus-finance/src/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock := database.NewMockDB()
	userRepo := NewUserRepository(db)
	expectedSQL := "INSERT INTO \"users\" (.+) VALUES (.+)"

	t.Run("create_user_successful", func(mt *testing.T) {
		addRow := mock.NewRows([]string{"id"}).AddRow("1")
		mock.ExpectBegin()
		mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
		mock.ExpectCommit()

		userDomain, err := userRepo.CreateUserRepository(models.NewUserDomain("TESTE", "TESTE@TESTE.COM", "123456"))

		assert.Nil(mt, err)
		assert.NotNil(mt, userDomain)
	})

	t.Run("create_user_no_rows_affected", func(mt *testing.T) {
		addRow := mock.NewRows([]string{"id"})
		mock.ExpectBegin()
		mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
		mock.ExpectCommit()

		userDomain, err := userRepo.CreateUserRepository(models.NewUserDomain("", "", ""))

		assert.NotNil(mt, err)
		assert.Nil(mt, userDomain)
	})
}
