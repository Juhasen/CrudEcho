package task

import (
	"RestCrud/internal/common"
	"RestCrud/internal/user"
	"github.com/google/uuid"
	"strings"
)

type Service struct {
	Repo     Repository
	UserRepo user.Repository
}

func NewService(repo Repository, userRepo user.Repository) *Service {
	return &Service{Repo: repo, UserRepo: userRepo}
}

func (s *Service) CreateTask(task *RequestDTO) error {
	if task.Title == "" || task.Description == "" || task.DueDate == "" || task.UserId.String() == "" || task.Status == "" {
		return ErrAllArgumentsRequired
	}

	if _, err := uuid.Parse(task.UserId.String()); err != nil {
		return ErrInvalidUserId
	}

	task.Status = common.Status(strings.ToLower(string(task.Status)))

	if err := task.Validate(); err != nil {
		return err
	}

	if _, err := s.UserRepo.FindByID(task.UserId.String()); err != nil {
		return ErrUserWithGivenIdDoesNotExist
	}

	return s.Repo.Save(TaskFromDTO(task))
}

func (s *Service) GetTaskByID(id string) (*ResponseDTO, error) {
	if id == "" {
		return nil, ErrTaskIdCannotBeEmpty
	}
	if _, err := uuid.Parse(id); err != nil {
		return nil, ErrIdIsNotValid
	}
	task, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return ToResponseDTO(task), nil
}

func (s *Service) GetAllTasks() ([]ResponseDTO, error) {
	tasks, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	var tasksDTO = make([]ResponseDTO, 0, len(tasks))
	for _, task := range tasks {
		tasksDTO = append(tasksDTO, *ToResponseDTO(&task))
	}
	return tasksDTO, err
}

func (s *Service) UpdateTask(id string, taskRequest *RequestDTO) error {
	if id == "" {
		return ErrTaskIdCannotBeEmpty
	}

	if _, err := uuid.Parse(id); err != nil {
		return ErrIdIsNotValid
	}

	if taskRequest.Status == "" &&
		taskRequest.Title == "" &&
		taskRequest.Description == "" &&
		taskRequest.DueDate == "" &&
		taskRequest.UserId == uuid.Nil {
		return ErrAtLeastOneFieldRequired
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
		task.DueDate, err = ParseDateStringToTime(taskRequest.DueDate)
		if err != nil {
			return ErrInvalidDateFormat
		}
	}

	if taskRequest.UserId != uuid.Nil {
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
			return ErrInvalidStatus
		}

		task.Status = taskRequest.Status
	}

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return ErrInvalidUserId
	}
	taskRequest.UserId = parsedUUID

	return s.Repo.Save(task)
}

func (s *Service) DeleteTask(id string) error {
	if id == "" {
		return ErrTaskIdCannotBeEmpty
	}
	if _, err := uuid.Parse(id); err != nil {
		return ErrIdIsNotValid
	}
	return s.Repo.Delete(id)
}
