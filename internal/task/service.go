package task

import (
	"RestCrud/internal/task/common"
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
	if task.Title == "" || task.Description == "" || task.DueDate == "" || task.UserId.String() == "" || task.Status == "" {
		return errors.ErrAllArgumentsRequired
	}

	if len(task.UserId.String()) < 36 {
		return errors.ErrInvalidUserId
	}

	task.Status = common.Status(strings.ToLower(string(task.Status)))

	if err := task.Validate(); err != nil {
		return err
	}

	// Check if the user ID exists
	if _, err := s.UserRepo.FindByID(task.UserId.String()); err != nil {
		return errors.ErrUserWithGivenIdDoesNotExist
	}

	return s.Repo.Save(model.TaskFromDTO(task))
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
		tasksDTO = append(tasksDTO, *task.ToResponseDTO())
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
	if taskRequest.DueDate != "" {
		if err := taskRequest.ValidateDate(); err != nil {
			return err
		}
		task.DueDate, err = common.ParseDateStringToTime(taskRequest.DueDate)
		if err != nil {
			return errors.ErrInvalidDateFormat
		}
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
		taskRequest.Status = common.Status(strings.ToLower(string(taskRequest.Status)))

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
