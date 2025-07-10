package task

import (
	"RestCrud/internal/common"
	"RestCrud/internal/model"
	"github.com/google/uuid"
	"time"
)

type ResponseDTO struct {
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	DueDate     string        `json:"due_date,omitempty"`
	UserId      uuid.UUID     `json:"user_id,omitempty"`
	Status      common.Status `json:"status" validate:"omitempty, oneof=pending in_progress completed cancelled"`
}

type RequestDTO struct {
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	DueDate     string        `json:"due_date,omitempty"`
	UserId      uuid.UUID     `json:"user_id,omitempty"`
	Status      common.Status `json:"status,omitempty" validate:"oneof=pending in_progress completed cancelled"`
}

func (t *RequestDTO) ValidateStatus() error {
	if t.Status != common.Pending &&
		t.Status != common.InProgress &&
		t.Status != common.Completed &&
		t.Status != common.Cancelled {
		return ErrInvalidStatus
	}
	return nil
}

func (t *RequestDTO) ValidateDate() error {
	// Parse due date
	parsedDate, err := time.Parse("02/01/2006", t.DueDate)
	if err != nil {
		return ErrInvalidDateFormat
	}

	// Check if the date is not in the past
	if parsedDate.Before(time.Now()) {
		return ErrDueDateInPast
	}

	return nil
}

func (t *RequestDTO) Validate() error {
	if err := t.ValidateDate(); err != nil {
		return err
	}
	if err := t.ValidateStatus(); err != nil {
		return err
	}
	return nil
}

func ToResponseDTO(t *model.Task) *ResponseDTO {
	return &ResponseDTO{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     FormatTimeToDateString(t.DueDate),
		UserId:      t.UserID,
		Status:      t.Status,
	}
}

func TaskFromDTO(dto *RequestDTO) *model.Task {
	convertedTime, err := ParseDateStringToTime(dto.DueDate)
	if err != nil {
		return nil
	}
	return &model.Task{
		Title:       dto.Title,
		Description: dto.Description,
		DueDate:     convertedTime,
		Status:      dto.Status,
		UserID:      dto.UserId,
	}
}
