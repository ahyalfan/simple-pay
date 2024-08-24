package dto

type UserRegisterReg struct {
	FullName string `json:"full_name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
