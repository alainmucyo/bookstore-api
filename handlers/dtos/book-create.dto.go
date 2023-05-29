package dtos

import (
	"bookstore/store/entities"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

func init() {
	validate = validator.New()
}

type BookCreateDTO struct {
	Title string `json:"title" validate:"required,min=2"`

	AuthorID string `json:"author_id" validate:"required"`

	Published time.Time `json:"published" validate:"required"`

	Price float64 `json:"price" validate:"required,gte=0"`

	ImageUrl string `json:"image_url" validate:"required,url"`
}

func (b *BookCreateDTO) Validate() error {
	err := validate.Struct(b)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.StructField()
			tag := err.Tag()
			switch field {
			case "Title":
				if tag == "required" {
					return fmt.Errorf("title is required")
				} else if tag == "min" {
					return fmt.Errorf("title should have a minimum length of 2")
				}
			case "AuthorID":
				if tag == "required" {
					return fmt.Errorf("author ID is required")
				}
			case "Published":
				if tag == "required" {
					return fmt.Errorf("published date is required")
				}
			case "Price":
				if tag == "required" {
					return fmt.Errorf("price is required")
				} else if tag == "gte" {
					return fmt.Errorf("price should be greater than or equal to 0")
				}
			case "ImageUrl":
				if tag == "required" {
					return fmt.Errorf("image URL is required")
				} else if tag == "url" {
					return fmt.Errorf("image URL is not valid")
				}
			}
		}
	}
	return nil
}

func (b *BookCreateDTO) ToEntity() entities.Book {
	authorIdUUID, _ := uuid.Parse(b.AuthorID)
	return entities.Book{
		Title:     b.Title,
		AuthorID:  authorIdUUID,
		Published: b.Published,
		Price:     b.Price,
	}
}
