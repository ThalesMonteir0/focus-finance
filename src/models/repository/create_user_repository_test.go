package repository

import (
	"focus-finance/src/configuration/database"
	"focus-finance/src/models"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock := database.NewMockDB()
	userRepo := NewUserRepository(db)

	addRow := mock.NewRows([]string{"id"}).AddRow("1")
	expectedSQL := "INSERT INTO \"users\" (.+) VALUES (.+)"
	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
	mock.ExpectCommit()

	userDomain, err := userRepo.CreateUserRepository(models.NewUserDomain("TESTE", "TESTE@TESTE.COM", "123456"))
	if err != nil {
		t.Fatalf("Failed to insert user %v", err)
	}

	if userDomain.GetID() == 0 {
		t.Fatalf("Failed to insert user")
	}
}
