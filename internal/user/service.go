package user

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateUser(user *User) error {
	if user.Name == "" {
		return ErrUserNameRequired
	}

	if user.Email == "" {
		return ErrUserEmailRequired
	}

	if user, _ := s.GetUserByID(user.ID); user == nil {
		return ErrUserAlreadyExists
	}

	return s.Repo.Save(user)
}

func (s *Service) GetUserByID(id string) (*User, error) {
	return s.Repo.FindByID(id)
}

func (s *Service) GetAllUsers() (*map[string]User, error) {
	return s.Repo.FindAll()
}

func (s *Service) UpdateUser(user *User) error {
	return s.Repo.Save(user)
}

func (s *Service) DeleteUser(id string) error {
	return s.Repo.Delete(id)
}
