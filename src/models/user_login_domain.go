package models

type userLoginDomain struct {
	email    string
	password string
}

func (u *userLoginDomain) GetEmail() string {
	return u.email
}

func (u *userLoginDomain) GetPassword() string {
	return u.password
}
