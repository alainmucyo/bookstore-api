package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Author struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	FirstName string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Bio       string    `gorm:"type:text" json:"bio"`
	Books     []Book    `gorm:"foreignKey:AuthorID" json:"books"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (*Author) TableName() string {
	return "authors"
}

func (author *Author) BeforeCreate(tx *gorm.DB) (err error) {
	author.Id = uuid.New()
	return
}
