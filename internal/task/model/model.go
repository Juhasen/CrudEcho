package model

import (
	"RestCrud/internal/task/common"
	"RestCrud/internal/task/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model

	ID          uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string        `gorm:"size:255;not null" json:"title"`
	Description string        `gorm:"type:text" json:"description"`
	DueDate     time.Time     `json:"due_date"`
	Status      common.Status `gorm:"size:100;not null" json:"status"`
	UserID      uuid.UUID     `gorm:"type:uuid;not null" json:"user_id"`
}

func (t *Task) ToResponseDTO() *dto.TaskResponseDTO {
	return &dto.TaskResponseDTO{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     common.FormatTimeToDateString(t.DueDate),
		UserId:      t.UserID,
		Status:      t.Status,
	}
}

func TaskFromDTO(dto *dto.TaskRequestDTO) *Task {
	convertedTime, err := common.ParseDateStringToTime(dto.DueDate)
	if err != nil {
		return nil
	}
	return &Task{
		Title:       dto.Title,
		Description: dto.Description,
		DueDate:     convertedTime,
		Status:      dto.Status,
		UserID:      dto.UserId,
	}
}
