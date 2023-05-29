package dtos

import (
	"bookstore/store/entities"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

type AuthorCreateDTO struct {
	FirstName string `json:"first_name" validate:"required,min=2"`

	LastName string `json:"last_name" validate:"required,min=2"`

	Bio string `json:"bio" validate:"required"`
}

func (a *AuthorCreateDTO) Validate() error {
	err := validate.Struct(a)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.StructField()
			tag := err.Tag()
			switch field {
			case "FirstName":
				if tag == "required" {
					return fmt.Errorf("first name is required")
				} else if tag == "min" {
					return fmt.Errorf("first name should have a minimum length of 2")
				}
			case "LastName":
				if tag == "required" {
					return fmt.Errorf("last name is required")
				} else if tag == "min" {
					return fmt.Errorf("last name should have a minimum length of 2")
				}
			case "Bio":
				if tag == "required" {
					return fmt.Errorf("bio is required")
				}
			}
		}
	}
	return nil
}

func (a *AuthorCreateDTO) ToAuthorEntity() entities.Author {
	return entities.Author{
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Bio:       a.Bio,
	}
}
