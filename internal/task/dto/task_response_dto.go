package dto

import (
	"RestCrud/internal/task/common"
	"github.com/google/uuid"
)

type TaskResponseDTO struct {
	Title       string        `json:"title, omitempty"`
	Description string        `json:"description, omitempty"`
	DueDate     string        `json:"due_date, omitempty"`
	UserId      uuid.UUID     `json:"user_id, omitempty"`
	Status      common.Status `json:"status" validate:"omitempty, oneof=pending in_progress completed cancelled"`
}
