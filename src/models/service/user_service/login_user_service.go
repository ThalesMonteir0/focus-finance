package user_service

import (
	"focus-finance/src/configuration/rest_err"
	"focus-finance/src/models"
	"golang.org/x/crypto/bcrypt"
)

func (us *userService) LoginUser(userDomain models.UserLoginDomainInterface) (string, *rest_err.RestErr) {
	user, err := us.repository.FindUserByEmail(userDomain.GetEmail())
	if err != nil {
		return "", err
	}

	ok := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(userDomain.GetPassword()))
	if ok != nil {
		return "", rest_err.NewBadRequestError("Password incorrect")
	}

	tokenValue, err := user.GenerateToken()
	if err != nil {
		return "", err
	}

	return tokenValue, nil
}
