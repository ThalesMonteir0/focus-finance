package models

type UserDomainInterface interface {
	GetID() int
	SetID(int)
	GetName() string
	GetEmail() string
	GetPassword() string
	GetConfirmPassword() string
}

func NewUserDomain(name, email, password, confirmPassword string) UserDomainInterface {
	return &userDomain{
		name:            name,
		email:           email,
		password:        password,
		confirmPassword: confirmPassword,
	}
}
