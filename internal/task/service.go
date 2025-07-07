package task

type Service struct {
	Repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateTask(task *Task) error {
	return s.Repo.Create(task)
}

func (s *Service) GetTaskByID(id string) (*Task, error) {
	return s.Repo.FindByID(id)
}

func (s *Service) GetAllTasks() ([]*Task, error) {
	return s.Repo.FindAll()
}

func (s *Service) UpdateTask(task *Task) error {
	return s.Repo.Update(task)
}

func (s *Service) DeleteTask(id string) error {
	return s.Repo.Delete(id)
}
