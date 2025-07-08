package task

import (
	"RestCrud/internal/task/dto"
	"RestCrud/internal/task/errors"
	"RestCrud/internal/task/model"
	"RestCrud/internal/user"
	"strings"
)

type Service struct {
	Repo     Repository
	UserRepo user.Repository
}

func NewService(repo Repository, userRepo user.Repository) *Service {
	return &Service{Repo: repo, UserRepo: userRepo}
}

func (s *Service) CreateTask(task *dto.TaskRequestDTO) error {
	if task.Title == "" || task.Description == "" || task.DueDate.String() == "" || task.UserId.String() == "" || task.Status == "" {
		return errors.ErrAllArgumentsRequired
	}

	task.Status = strings.ToLower(task.Status)

	if err := task.Validate(); err != nil {
		return err
	}

	// Check if the user ID exists
	if _, err := s.UserRepo.FindByID(task.UserId.String()); err != nil {
		return err
	}

	return s.Repo.Save(dtoToTask(task))
}

func (s *Service) GetTaskByID(id string) (*model.Task, error) {
	if id == "" {
		return nil, errors.ErrTaskIdCannotBeEmpty
	}
	return s.Repo.FindByID(id)
}

func (s *Service) GetAllTasks() ([]dto.TaskResponseDTO, error) {
	tasks, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	var tasksDTO = make([]dto.TaskResponseDTO, 0, len(tasks))
	for _, task := range tasks {
		tasksDTO = append(tasksDTO, *taskToDTO(&task))
	}
	return tasksDTO, err
}

func (s *Service) UpdateTask(id string, taskRequest *dto.TaskRequestDTO) error {
	if id == "" {
		return errors.ErrTaskIdCannotBeEmpty
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
	if taskRequest.DueDate.String() != "" {
		if err := taskRequest.ValidateDate(); err != nil {
			return err
		}
		task.DueDate = taskRequest.DueDate
	}
	if taskRequest.UserId.String() != "" {
		// Check if the user ID exists
		if _, err := s.UserRepo.FindByID(taskRequest.UserId.String()); err != nil {
			return err
		}
		task.UserID = taskRequest.UserId
	}
	if taskRequest.Status != "" {
		// Normalize the status to lowercase
		taskRequest.Status = strings.ToLower(taskRequest.Status)

		if err := taskRequest.ValidateStatus(); err != nil {
			return errors.ErrInvalidStatus
		}

		task.Status = taskRequest.Status
	}

	return s.Repo.Save(task)
}

func (s *Service) DeleteTask(id string) error {
	if id == "" {
		return errors.ErrTaskIdCannotBeEmpty
	}
	return s.Repo.Delete(id)
}

func taskToDTO(t *model.Task) *dto.TaskResponseDTO {
	if t == nil {
		return nil
	}
	return &dto.TaskResponseDTO{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		UserId:      t.UserID,
		Status:      t.Status,
	}
}

func dtoToTask(t *dto.TaskRequestDTO) *model.Task {
	if t == nil {
		return nil
	}
	return &model.Task{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		UserID:      t.UserId,
		Status:      t.Status,
	}
}
