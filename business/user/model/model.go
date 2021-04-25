package model

type UserModel struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type UserLoginModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
