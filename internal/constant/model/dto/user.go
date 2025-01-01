package dto

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required.Error("email is required"), is.Email.Error("email is not valid")),
		validation.Field(&u.Password, validation.Required.Error("password is required"),
			validation.Length(8, 20).Error("password must be at least 8 characters long")),
	)
}
