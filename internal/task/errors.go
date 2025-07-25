package task

import "errors"

var (
	ErrTaskIdCannotBeEmpty         = errors.New("task ID cannot be empty")
	ErrTaskWithGivenIdNotFound     = errors.New("task with given ID not found")
	ErrNoTasksFound                = errors.New("no tasks found")
	ErrAllArgumentsRequired        = errors.New("all arguments are required to create a task")
	ErrInvalidStatus               = errors.New("status must be one of Pending, InProgress, Completed or Cancelled")
	ErrInvalidDateFormat           = errors.New("invalid date format, expected DD/MM/YYYY")
	ErrDueDateInPast               = errors.New("due date must be in the future")
	ErrFailedToDeleteTask          = errors.New("failed to delete task")
	ErrLoadDataFailed              = errors.New("failed to load task data")
	ErrSaveDataFailed              = errors.New("failed to save data")
	ErrUserWithGivenIdDoesNotExist = errors.New("user with given ID does not exist")
	ErrInvalidUserId               = errors.New("invalid user ID, must be a valid UUID")
	ErrIdIsNotValid                = errors.New("ID is not valid UUID format")
	ErrAtLeastOneFieldRequired     = errors.New("at least one field (status, title, description, due date, user ID) is required for update")
)
