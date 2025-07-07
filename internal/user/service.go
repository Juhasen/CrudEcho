package user

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateUser(user *User) error {
	return s.Repo.Create(user)
}

func (s *Service) GetUserByID(id string) (*User, error) {
	return s.Repo.FindByID(id)
}

func (s *Service) GetAllUsers() ([]*User, error) {
	return s.Repo.FindAll()
}

func (s *Service) UpdateUser(user *User) error {
	return s.Repo.Update(user)
}

func (s *Service) DeleteUser(id string) error {
	return s.Repo.Delete(id)
}
