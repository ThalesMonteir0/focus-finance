package models

type userDomain struct {
	id              int
	name            string
	email           string
	password        string
	confirmPassword string
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
func (ud *userDomain) GetConfirmPassword() string {
	return ud.confirmPassword
}
