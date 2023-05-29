package dtos

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// swagger:model User
type RegisterDTO struct {
	// Name must be at least 2 characters long.
	Name string `json:"name" validate:"required,min=2"`

	// Email must be a valid email string of the form "bob@example.com".
	Email string `json:"email" validate:"required,email"`

	// Password must be at least 6 characters long.
	Password string `json:"password" validate:"required,min=6"`
}

func (u *RegisterDTO) Validate() error {
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.StructField() // field name
			tag := err.Tag()           // tag that failed
			switch field {
			case "Name":
				if tag == "required" {
					return fmt.Errorf("name is required")
				} else if tag == "min" {
					return fmt.Errorf("name should have a minimum length of 2")
				}
			case "Email":
				if tag == "required" {
					return fmt.Errorf("email is required")
				} else if tag == "email" {
					return fmt.Errorf("please provide a valid email")
				}
			case "Password":
				if tag == "required" {
					return fmt.Errorf("password is required")
				} else if tag == "min" {
					return fmt.Errorf("password should have a minimum length of 6")
				}
			}
		}
	}
	return nil
}
