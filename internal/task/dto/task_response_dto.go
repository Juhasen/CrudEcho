package dto

import (
	"RestCrud/internal/task/common"
	"github.com/google/uuid"
)

type TaskResponseDTO struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	DueDate     string        `json:"due_date"`
	UserId      uuid.UUID     `json:"user_id"`
	Status      common.Status `json:"status" validate:"oneof=pending in_progress completed cancelled"`
}
