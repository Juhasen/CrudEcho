package task

import (
	"RestCrud/internal/user"
	generated "RestCrud/openapi"
	"github.com/google/uuid"
)

type Service struct {
	Repo     Repository
	UserRepo user.Repository
}

func NewService(repo Repository, userRepo user.Repository) *Service {
	return &Service{Repo: repo, UserRepo: userRepo}
}

func (s *Service) CreateTask(task *generated.TaskRequest) error {
	if task.Title == "" || task.Description == "" || task.UserId.String() == "" {
		return ErrAllArgumentsRequired
	}

	if _, err := uuid.Parse(task.UserId.String()); err != nil {
		return ErrInvalidUserId
	}

	if err := Validate(task); err != nil {
		return err
	}

	if _, err := s.UserRepo.FindByID(task.UserId.String()); err != nil {
		return ErrUserWithGivenIdDoesNotExist
	}

	return s.Repo.Save(TaskFromDTO(task))
}

func (s *Service) GetTaskByID(id string) (*generated.TaskResponse, error) {
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

func (s *Service) GetAllTasks() ([]generated.TaskResponse, error) {
	tasks, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	var tasksDTO = make([]generated.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		tasksDTO = append(tasksDTO, *ToResponseDTO(&task))
	}
	return tasksDTO, err
}

func (s *Service) UpdateTask(id string, taskRequest *generated.TaskRequest) error {
	if id == "" {
		return ErrTaskIdCannotBeEmpty
	}

	if _, err := uuid.Parse(id); err != nil {
		return ErrIdIsNotValid
	}

	if taskRequest.Title == "" &&
		taskRequest.Description == "" &&
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

	if err := ValidateDate(taskRequest); err != nil {
		return err
	}
	task.DueDate, err = ParseDateStringToTime(taskRequest.DueDate)
	if err != nil {
		return ErrInvalidDateFormat
	}

	if taskRequest.UserId != uuid.Nil {
		// Check if the user ID exists
		if _, err := s.UserRepo.FindByID(taskRequest.UserId.String()); err != nil {
			return err
		}

		task.UserID = taskRequest.UserId
	}

	if err := ValidateStatus(taskRequest); err != nil {
		return ErrInvalidStatus
	}

	task.Status = *taskRequest.Status

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
