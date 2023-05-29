package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	AuthorID    uuid.UUID `gorm:"type:uuid;not null;column:author_id" json:"-"`
	Author      Author    `json:"author,omitempty" gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Published   time.Time `gorm:"not null" json:"published"`
	Price       float64   `gorm:"type:decimal(10,2);not null;default:0.0" json:"price"`
	ImageUrl    string    `gorm:"type:varchar(255)" json:"image_url"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (*Book) TableName() string {
	return "books"
}

func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	book.Id = uuid.New()
	return
}
