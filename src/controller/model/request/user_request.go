package request

type UserRequest struct {
	Name            string `json:"name" binding:"required,min=2"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" biding:"required,email"`
	Password string `json:"password" biding:"required,min=6"`
}
