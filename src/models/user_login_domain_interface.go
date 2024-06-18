package models

type UserLoginDomainInterface interface {
	GetEmail() string
	GetPassword() string
}

func NewUserLoginDomain(email, password string) UserLoginDomainInterface {
	return &userLoginDomain{
		password: password,
		email:    email,
	}
}
