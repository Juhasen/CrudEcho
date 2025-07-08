package dto

import (
	"github.com/google/uuid"
	"time"
)

type TaskResponseDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	UserId      uuid.UUID `json:"user_id"`
	Status      string    `json:"status" validate:"oneof=pending in_progress completed cancelled"`
}
