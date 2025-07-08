package user

import (
	"RestCrud/internal/user/dto"
	"RestCrud/internal/user/model"
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

	if _, err := s.Repo.FindByEmail(user.Email); err == nil {
		return ErrUserAlreadyExists
	}

	return s.Repo.Save(user)
}

func (s *Service) GetUserByID(id string) (*dto.UserResponseDTO, error) {
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

	existingUser, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if _, err := s.Repo.FindByEmail(existingUser.Email); err == nil {
		return ErrUserAlreadyExists
	}

	return s.Repo.Save(userToDTO(existingUser))
}

func (s *Service) DeleteUser(id string) error {
	if id == "" {
		return ErrUserIDRequired
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
