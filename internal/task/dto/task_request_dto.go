package dto

import (
	"RestCrud/internal/task/errors"
	"RestCrud/internal/task/model"
	"github.com/google/uuid"
	"time"
)

type TaskRequestDTO struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"due_date,omitempty"`
	UserId      uuid.UUID `json:"user_id,omitempty"`
	Status      string    `json:"status,omitempty" validate:"oneof=pending in_progress completed cancelled"`
}

func (t *TaskRequestDTO) ValidateStatus() error {
	if t.Status != string(model.StatusPending) &&
		t.Status != string(model.StatusInProgress) &&
		t.Status != string(model.StatusCompleted) &&
		t.Status != string(model.StatusCancelled) {
		return errors.ErrInvalidStatus
	}
	return nil
}

func (t *TaskRequestDTO) ValidateDate() error {
	// Parse due date
	parsedDate, err := time.Parse("02/01/2006", t.DueDate.String())
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

func (t *TaskRequestDTO) ToModel() *model.Task {
	return &model.Task{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		UserID:      t.UserId,
		Status:      t.Status,
	}
}
