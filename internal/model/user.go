package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name  string    `gorm:"size:255;not null" json:"name"`
	Email string    `gorm:"size:255;unique;not null" json:"email"`
	Tasks []Task    `gorm:"foreignKey:UserID"`
}
