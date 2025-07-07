package task

type Repository interface {
	Save(task *Task) error
	FindByID(id string) (*Task, error)
	FindAll() ([]*Task, error)
	Delete(id string) error
}

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Save(task *Task) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) FindByID(id string) (*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) FindAll() ([]*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
