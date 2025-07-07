package task

import "RestCrud/pkg/repository"

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Create(task *Task) error {
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

func (r *Repo) Update(task *Task) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repo) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
