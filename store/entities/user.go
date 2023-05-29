package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Status    string    `json:"status" gorm:"column:status;default:active"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (*User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	user.Id = id
	return err
}
