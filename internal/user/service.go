package user

import (
	"RestCrud/internal/model"
	"RestCrud/kafka"
	generated "RestCrud/openapi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strings"
)

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateUser(user *generated.UserResponse) (*model.User, error) {
	if user.Name == "" {
		return nil, ErrUserNameRequired
	}

	if user.Email == "" {
		return nil, ErrUserEmailRequired
	}

	if !strings.Contains(string(user.Email), "@") {
		return nil, ErrUserEmailInvalid
	}

	if user, _ := s.Repo.FindByEmail(string(user.Email)); user != nil {
		return nil, ErrUserAlreadyExists
	}

	savedUser := dtoToUser(user)

	if err := s.Repo.Save(savedUser); err != nil {
		return nil, err
	}

	if err := kafka.ProduceTodoEvent(savedUser, kafka.CREATE, savedUser.ID.String()); err != nil {
		return nil, errors.Wrap(err, "kafka failed to produce message")
	}

	return savedUser, nil
}

func (s *Service) GetUserByID(id string) (*generated.UserResponse, error) {
	if id == "" {
		return nil, ErrUserIDRequired
	}

	if _, err := uuid.Parse(id); err != nil {
		return nil, ErrIdIsNotValid
	}

	user, err := s.Repo.FindByID(id)
	return userToDTO(user), err
}

func (s *Service) GetAllUsers() ([]generated.UserResponse, error) {
	users, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var usersDTO = make([]generated.UserResponse, 0, len(users))
	for _, user := range users {
		usersDTO = append(usersDTO, *userToDTO(&user))
	}

	return usersDTO, err
}

func (s *Service) UpdateUser(id string, user *generated.UserRequest) error {
	if id == "" {
		return ErrUserIDRequired
	}

	if _, err := uuid.Parse(id); err != nil {
		return ErrIdIsNotValid
	}

	if user.Name == "" && user.Email == "" {
		return ErrAtLeastOneFieldRequired
	}

	existingUser, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	if user.Email != "" {
		if user, _ := s.Repo.FindByEmail(string(user.Email)); user != nil {
			return ErrUserAlreadyExists
		}
		existingUser.Email = string(user.Email)
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}

	if err := kafka.ProduceTodoEvent(existingUser, kafka.EDIT, existingUser.ID.String()); err != nil {
		return errors.Wrap(err, "kafka failed to produce message")
	}

	return s.Repo.Save(existingUser)
}

func (s *Service) DeleteUser(id string) error {
	if id == "" {
		return ErrUserIDRequired
	}

	if _, err := uuid.Parse(id); err != nil {
		return ErrIdIsNotValid
	}

	return s.Repo.Delete(id)
}
