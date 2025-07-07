package user

type Repository interface {
	Create(user *User) error
	FindByID(id string) (*User, error)
	FindAll() ([]*User, error)
	Update(user *User) error
	Delete(id string) error
}

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Create(user *User) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) FindByID(id string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) FindAll() ([]*User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) Update(user *User) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
