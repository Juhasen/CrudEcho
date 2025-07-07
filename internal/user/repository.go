package user

import "RestCrud/pkg/repository"

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
