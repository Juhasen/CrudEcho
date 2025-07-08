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

func (s *Service) CreateUser(user *dto.UserDTO) error {
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

func (s *Service) GetUserByID(id string) (*dto.UserDTO, error) {
	user, err := s.Repo.FindByID(id)
	return userToDTO(user), err
}

func (s *Service) GetAllUsers() (*map[string]dto.UserDTO, error) {

	users, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	var usersDTO = make(map[string]dto.UserDTO)
	for _, user := range *users {
		dtoUser := userToDTO(&user)
		if dtoUser != nil {
			usersDTO[user.ID.String()] = *dtoUser
		}
	}

	return &usersDTO, err
}

func (s *Service) UpdateUser(id string, user *dto.UserUpdateDTO) error {
	return s.Repo.Update(id, user)
}

func (s *Service) DeleteUser(id string) error {
	return s.Repo.Delete(id)
}

func userToDTO(u *model.User) *dto.UserDTO {
	if u == nil {
		return nil
	}
	return &dto.UserDTO{
		Name:  u.Name,
		Email: u.Email,
	}
}
