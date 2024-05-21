package repository

import (
	"focus-finance/src/configuration/database"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserRepository_FindUserByID(t *testing.T) {
	db, mock := database.NewMockDB()
	userRepo := NewUserRepository(db)
	expectedSQL := "SELECT \\* FROM \"users\" WHERE \"users\"\\.\"id\" = \\$1 AND \"users\"\\.\"deleted_at\" IS NULL ORDER BY \"users\"\\.\"id\" LIMIT \\$2"

	t.Run("find_user_by_id_successful", func(mt *testing.T) {
		rows := mock.NewRows([]string{"id", "name", "email", "password", "deleted_at", "updated_at"}).
			AddRow(1, "TESTE", "TESTE@TESTE.COM", "123456", time.Now(), time.Now())
		mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

		userDomain, err := userRepo.FindUserByID(1)
		assert.Nil(mt, err)
		assert.NotNil(mt, userDomain)
	})

	t.Run("find_user_by_id_error", func(mt *testing.T) {
		rows := mock.NewRows([]string{"id", "name", "email", "password", "deleted_at", "updated_at"})
		mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

		userDomain, err := userRepo.FindUserByID(1)

		assert.NotNil(mt, err)
		assert.Nil(mt, userDomain)
	})

}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	db, mock := database.NewMockDB()
	userRepo := NewUserRepository(db)
	expectedSQL := "SELECT \\* FROM \"users\" WHERE email = \\$1 AND \"users\"\\.\"deleted_at\" IS NULL ORDER BY \"users\"\\.\"id\" LIMIT \\$2"

	t.Run("find_user_by_email_successful", func(mt *testing.T) {
		rows := mock.NewRows([]string{"id", "name", "email", "password", "deleted_at", "updated_at"}).AddRow(1, "TESTE", "TESTE@TESTE.COM", "123456", time.Now(), time.Now())
		mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

		userDomain, err := userRepo.FindUserByEmail("TESTE@TESTE.COM")

		assert.Nil(mt, err)
		assert.NotNil(mt, userDomain)
		assert.Equal(mt, userDomain.GetEmail(), "TESTE@TESTE.COM")
	})

	t.Run("find_user_by_email_error", func(mt *testing.T) {
		rows := mock.NewRows([]string{"id", "name", "email", "password", "deleted_at", "updated_at"})
		mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

		userDomain, err := userRepo.FindUserByEmail("TESTE@TESTE.COM")

		assert.NotNil(mt, err)
		assert.Nil(mt, userDomain)
	})

}
