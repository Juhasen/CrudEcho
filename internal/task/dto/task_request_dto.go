package dto

import (
	"RestCrud/internal/task/common"
	"RestCrud/internal/task/errors"
	"github.com/google/uuid"
	"time"
)

type TaskRequestDTO struct {
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	DueDate     string        `json:"due_date,omitempty"`
	UserId      uuid.UUID     `json:"user_id,omitempty"`
	Status      common.Status `json:"status,omitempty" validate:"oneof=pending in_progress completed cancelled"`
}

func (t *TaskRequestDTO) ValidateStatus() error {
	if t.Status != common.Pending &&
		t.Status != common.InProgress &&
		t.Status != common.Completed &&
		t.Status != common.Cancelled {
		return errors.ErrInvalidStatus
	}
	return nil
}

func (t *TaskRequestDTO) ValidateDate() error {
	// Parse due date
	parsedDate, err := time.Parse("02/01/2006", t.DueDate)
	if err != nil {
		return errors.ErrInvalidDateFormat
	}

	// Check if the date is in the future
	if parsedDate.Before(time.Now()) {
		return errors.ErrDueDateInPast
	}

	return nil
}

func (t *TaskRequestDTO) Validate() error {

	if err := t.ValidateDate(); err != nil {
		return err
	}

	if err := t.ValidateStatus(); err != nil {
		return err
	}
	return nil
}
