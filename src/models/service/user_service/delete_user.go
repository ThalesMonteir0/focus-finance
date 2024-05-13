package user_service

import "focus-finance/src/configuration/rest_err"

func (us *userService) DeleteUser(userID int) *rest_err.RestErr {
	if err := us.repository.DeleteUserRepository(userID); err != nil {
		return err
	}

	return nil
}
