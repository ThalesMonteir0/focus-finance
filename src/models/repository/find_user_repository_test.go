package repository

import (
	"focus-finance/src/configuration/database"
	"strings"
	"testing"
	"time"
)

func TestUserRepository_FindUserByID(t *testing.T) {
	db, mock := database.NewMockDB()
	userRepo := NewUserRepository(db)

	expectedSQL := "SELECT \\* FROM \"users\" WHERE \"users\"\\.\"id\" = \\$1 AND \"users\"\\.\"deleted_at\" IS NULL ORDER BY \"users\"\\.\"id\" LIMIT \\$2"
	rows := mock.NewRows([]string{"id", "name", "email", "password", "deleted_at", "updated_at"}).
		AddRow(1, "TESTE", "TESTE@TESTE.COM", "123456", time.Now(), time.Now())
	mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	userDomain, err := userRepo.FindUserByID(1)
	if err != nil {
		t.Fatalf("unable get user err: %v", err)
	}

	if userDomain.GetID() == 0 {
		t.Fatalf("unable get user, userID = 0")
	}
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	db, mock := database.NewMockDB()
	userRepo := NewUserRepository(db)

	expectedSQL := "SELECT \\* FROM \"users\" WHERE email = \\$1 AND \"users\"\\.\"deleted_at\" IS NULL ORDER BY \"users\"\\.\"id\" LIMIT \\$2"
	rows := mock.NewRows([]string{"id", "name", "email", "password", "deleted_at", "updated_at"}).AddRow(1, "TESTE", "TESTE@TESTE.COM", "123456", time.Now(), time.Now())
	mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	userDomain, err := userRepo.FindUserByEmail("TESTE@TESTE.COM")
	if err != nil {
		t.Fatalf("unable get user by email err: %v", err)
	}

	if !strings.EqualFold(userDomain.GetEmail(), "TESTE@TESTE.COM") {
		t.Fatalf("unable get user by email, email is not equal")
	}
}
