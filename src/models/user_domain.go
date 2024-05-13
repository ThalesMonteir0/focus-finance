package models

import (
	"focus-finance/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type userDomain struct {
	id       int
	name     string
	email    string
	password string
}

func (ud *userDomain) GetID() int {
	return ud.id
}
func (ud *userDomain) SetID(id int) {
	ud.id = id
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) SetEmailToUppercase() {
	ud.email = strings.ToUpper(ud.email)
}
func (ud *userDomain) SetNameToUppercase() {
	ud.name = strings.ToUpper(ud.name)
}
func (ud *userDomain) EncryptPassword() *rest_err.RestErr {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(ud.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		err := rest_err.NewInternalServerError("unable encrypt password")
		return err
	}

	ud.password = string(hashPassword)

	return nil
}
