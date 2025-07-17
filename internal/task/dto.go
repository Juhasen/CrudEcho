package task

import (
	"RestCrud/internal/model"
	generated "RestCrud/openapi"
	"time"
)

func ValidateStatus(t *generated.TaskRequest) error {
	if *t.Status != generated.CANCELLED &&
		*t.Status != generated.COMPLETED &&
		*t.Status != generated.INPROGRESS &&
		*t.Status != generated.PENDING {
		return ErrInvalidStatus
	}
	return nil
}

func ValidateDate(t *generated.TaskRequest) error {
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

func Validate(t *generated.TaskRequest) error {
	if err := ValidateDate(t); err != nil {
		return err
	}
	if err := ValidateStatus(t); err != nil {
		return err
	}
	return nil
}

func ToResponseDTO(t *model.Task) *generated.TaskResponse {
	return &generated.TaskResponse{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		UserId:      t.UserID,
		Status:      t.Status,
	}
}

func TaskFromDTO(dto *generated.TaskRequest) *model.Task {
	convertedTime, err := ParseDateStringToTime(dto.DueDate)
	if err != nil {
		return nil
	}
	return &model.Task{
		Title:       dto.Title,
		Description: dto.Description,
		DueDate:     convertedTime,
		Status:      *dto.Status,
		UserID:      dto.UserId,
	}
}
