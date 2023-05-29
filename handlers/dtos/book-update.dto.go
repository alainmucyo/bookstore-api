package dtos

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

func init() {
	validate = validator.New()
}

type BookUpdateDTO struct {
	Title string `json:"title" validate:"min=2,omitempty"`

	AuthorID string `json:"author_id"`

	Published time.Time `json:"published"`

	Price float64 `json:"price" validate:"gte=0,omitempty"`

	ImageUrl string `json:"image_url" validate:"omitempty,url"`
}

func (b *BookUpdateDTO) Validate() error {
	err := validate.Struct(b)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.StructField()
			tag := err.Tag()
			switch field {
			case "Title":
				if tag == "min" {
					return fmt.Errorf("title should have a minimum length of 2")
				}
			case "Price":
				if tag == "gte" {
					return fmt.Errorf("price should be greater than or equal to 0")
				}
			case "ImageUrl":
				if tag == "url" {
					return fmt.Errorf("image URL is not valid")
				}
			}
		}
	}
	return nil
}
