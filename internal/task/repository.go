package task

import (
	"RestCrud/internal/db"
	"errors"
	"fmt"
)

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
	tasks := make(map[string]Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return err
	}

	if task.Id == "" {
		return errors.New("error: task ID cannot be empty")
	}

	if _, exists := tasks[task.Id]; exists {
		return fmt.Errorf("error: task with ID:%s already exists", task.Id)
	}

	tasks[task.Id] = *task

	if err := db.SaveData(db.TaskFile, tasks); err != nil {
		return err
	}

	return nil
}

func (r *Repo) FindByID(id string) (*Task, error) {
	tasks := make(map[string]Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return &Task{}, err
	}

	task, found := tasks[id]
	if !found {
		return &Task{}, fmt.Errorf("error: task with ID:%d not found", id)
	}

	return &task, nil
}

func (r *Repo) FindAll() (*map[string]Task, error) {
	tasks := make(map[string]Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("error: no tasks found")
	}

	return &tasks, nil
}

func (r *Repo) Delete(id string) error {
	tasks := make(map[string]Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return err
	}

	if _, found := tasks[id]; !found {
		return fmt.Errorf("error: task with ID:%d not found", id)
	}

	delete(tasks, id)

	if err := db.SaveData(db.TaskFile, tasks); err != nil {
		return err
	}

	return nil
}
