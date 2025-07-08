package task

import (
	"RestCrud/internal/user"
	"time"
)

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateTask(task *Task) error {
	return s.Repo.Save(task)
}

func (s *Service) GetTaskByID(id string) (*Task, error) {
	return s.Repo.FindByID(id)
}

func (s *Service) GetAllTasks() (*map[string]Task, error) {
	return s.Repo.FindAll()
}

func (s *Service) UpdateTask(task *Task) error {
	return s.Repo.Save(task)
}

func (s *Service) DeleteTask(id string) error {
	if id == "" {
		return ErrTaskIdCannotBeEmpty
	}
	return s.Repo.Delete(id)
}

func validateTask(task *Task, repository user.Repository) error {
	if task.Title == "" || task.Description == "" || task.DueDate == "" || task.UserId == "" || task.Status == "" {
		return ErrAllArgumentsRequired
	}

	// Check if the status is valid
	if task.Status != string(StatusPending) && task.Status != string(StatusInProgress) && task.Status != string(StatusCompleted) && task.Status != string(StatusCancelled) {
		return ErrInvalidStatus
	}

	// Check if the user ID exists
	if _, err := repository.FindByID(task.UserId); err != nil {
		return err
	}

	// Parse due date
	parsedDate, err := time.Parse("02/01/2006", task.DueDate)
	if err != nil {
		return ErrInvalidDateFormat
	}

	// Check if the date is in the future
	if parsedDate.Before(time.Now()) {
		return ErrDueDateInPast
	}

	return nil
}
