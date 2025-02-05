package createUser

type InputCreateUser struct {
	Username string `json:"username" validate:"required,alphanum,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"passwrod" validate:"required,min=6,max=20"`
}
