package task

import (
	"RestCrud/internal/task/dto"
	"RestCrud/internal/user"
	"strings"
	"time"
)

type Service struct {
	Repo     Repository
	UserRepo user.Repository
}

func NewService(repo Repository, userRepo user.Repository) *Service {
	return &Service{Repo: repo, UserRepo: userRepo}
}

func (s *Service) CreateTask(task *dto.TaskRequestDTO) error {
	if task.Title == "" || task.Description == "" || task.DueDate == "" || task.UserId == "" || task.Status == "" {
		return ErrAllArgumentsRequired
	}

	// Check if the status is valid
	if task.Status != string(StatusPending) && task.Status != string(StatusInProgress) && task.Status != string(StatusCompleted) && task.Status != string(StatusCancelled) {
		return ErrInvalidStatus
	}

	// Check if the user ID exists
	if _, err := s.UserRepo.FindByID(task.UserId); err != nil {
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
	return s.Repo.Save(dtoToTask(task))
}

func (s *Service) GetTaskByID(id string) (*Task, error) {
	if id == "" {
		return nil, ErrTaskIdCannotBeEmpty
	}
	return s.Repo.FindByID(id)
}

func (s *Service) GetAllTasks() (*map[string]dto.TaskResponseDTO, error) {
	tasks, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	var tasksDTO = make(map[string]dto.TaskResponseDTO)
	for _, task := range *tasks {
		dtoTask := taskToDTO(&task)
		if dtoTask != nil {
			tasksDTO[task.ID] = *dtoTask
		}
	}

	return &tasksDTO, err
}

func (s *Service) UpdateTask(id string, taskRequest *dto.TaskRequestDTO) error {
	if id == "" {
		return ErrTaskIdCannotBeEmpty
	}

	task, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	if taskRequest.Title != "" {
		task.Title = taskRequest.Title
	}
	if taskRequest.Description != "" {
		task.Description = taskRequest.Description
	}
	if taskRequest.DueDate != "" {
		// Validate the due date format
		parsedDate, err := time.Parse("02/01/2006", taskRequest.DueDate)
		if err != nil {
			return ErrInvalidDateFormat
		}
		// Check if the due date is in the past
		if parsedDate.Before(time.Now()) {
			return ErrDueDateInPast
		}
		task.DueDate = taskRequest.DueDate
	}
	if taskRequest.UserId != "" {
		// Check if the user ID exists
		if _, err := s.UserRepo.FindByID(taskRequest.UserId); err != nil {
			return err
		}
		task.UserId = taskRequest.UserId
	}
	if taskRequest.Status != "" {
		// Normalize the status to lowercase
		taskRequest.Status = strings.ToLower(taskRequest.Status)

		// Check if the status is valid
		if taskRequest.Status != string(StatusPending) && taskRequest.Status != string(StatusInProgress) && taskRequest.Status != string(StatusCompleted) && taskRequest.Status != string(StatusCancelled) {
			return ErrInvalidStatus
		}
		task.Status = taskRequest.Status
	}

	return s.Repo.Save(task)
}

func (s *Service) DeleteTask(id string) error {
	if id == "" {
		return ErrTaskIdCannotBeEmpty
	}
	return s.Repo.Delete(id)
}

func taskToDTO(t *Task) *dto.TaskResponseDTO {
	if t == nil {
		return nil
	}
	return &dto.TaskResponseDTO{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		UserId:      t.UserId,
		Status:      t.Status,
	}
}

func dtoToTask(t *dto.TaskRequestDTO) *Task {
	if t == nil {
		return nil
	}
	return &Task{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		UserId:      t.UserId,
		Status:      t.Status,
	}
}
