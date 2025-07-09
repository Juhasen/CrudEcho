package user

import (
	"RestCrud/internal/user/dto"
	"RestCrud/internal/user/model"
	"github.com/google/uuid"
	"strings"
)

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateUser(user *dto.UserResponseDTO) error {
	if user.Name == "" {
		return ErrUserNameRequired
	}

	if user.Email == "" {
		return ErrUserEmailRequired
	}

	if !strings.Contains(user.Email, "@") {
		return ErrUserEmailInvalid
	}

	if user, _ := s.Repo.FindByEmail(user.Email); user != nil {
		return ErrUserAlreadyExists
	}

	return s.Repo.Save(dtoToUser(user))
}

func (s *Service) GetUserByID(id string) (*dto.UserResponseDTO, error) {
	if id == "" {
		return nil, ErrUserIDRequired
	}

	if _, err := uuid.Parse(id); err != nil {
		return nil, ErrIdIsNotValid
	}

	user, err := s.Repo.FindByID(id)
	return userToDTO(user), err
}

func (s *Service) GetAllUsers() ([]dto.UserResponseDTO, error) {
	users, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var usersDTO = make([]dto.UserResponseDTO, 0, len(users))
	for _, user := range users {
		usersDTO = append(usersDTO, *userToDTO(&user))
	}

	return usersDTO, err
}

func (s *Service) UpdateUser(id string, user *dto.UserRequestDTO) error {
	if id == "" {
		return ErrUserIDRequired
	}
	if user.ID != id {
		return ErrUserIDMismatch
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
		if user, _ := s.Repo.FindByEmail(existingUser.Email); user != nil {
			return ErrUserAlreadyExists
		}
		existingUser.Email = user.Email
	}

	if user.Name != "" {
		existingUser.Name = user.Name
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

func userToDTO(u *model.User) *dto.UserResponseDTO {
	if u == nil {
		return nil
	}
	return &dto.UserResponseDTO{
		Name:  u.Name,
		Email: u.Email,
	}
}

func dtoToUser(u *dto.UserResponseDTO) *model.User {
	if u == nil {
		return nil
	}
	return &model.User{
		Name:  u.Name,
		Email: u.Email,
	}
}
