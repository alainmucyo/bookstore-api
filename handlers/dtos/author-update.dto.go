package dtos

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

type AuthorUpdateDTO struct {
	FirstName string `json:"first_name" validate:"omitempty,min=2"`

	LastName string `json:"last_name" validate:"omitempty,min=2"`

	Bio string `json:"bio" validate:"omitempty"`
}

func (a *AuthorUpdateDTO) Validate() error {
	err := validate.Struct(a)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.StructField()
			tag := err.Tag()
			switch field {
			case "FirstName":
				if tag == "min" {
					return fmt.Errorf("first name should have a minimum length of 2")
				}
			case "LastName":
				if tag == "min" {
					return fmt.Errorf("last name should have a minimum length of 2")
				}
			case "Bio":
			}
		}
	}
	return nil
}
