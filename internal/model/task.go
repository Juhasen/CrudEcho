package model

import (
	generated "RestCrud/openapi"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model

	ID          uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string           `gorm:"size:255;not null" json:"title"`
	Description string           `gorm:"type:text" json:"description"`
	DueDate     time.Time        `json:"due_date"`
	Status      generated.Status `gorm:"size:100;not null" json:"status"`
	UserID      uuid.UUID        `gorm:"type:uuid;not null" json:"user_id"`
}
