package models

import "focus-finance/src/configuration/rest_err"

type UserDomainInterface interface {
	GetID() int
	SetID(int)
	GetName() string
	GetEmail() string
	GetPassword() string
	SetEmailToUppercase()
	SetNameToUppercase()
	EncryptPassword() *rest_err.RestErr
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(name, email, password string) UserDomainInterface {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(name, email string) UserDomainInterface {
	return &userDomain{
		name:  name,
		email: email,
	}
}
