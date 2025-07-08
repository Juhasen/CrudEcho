package task

import (
	"RestCrud/internal/db"
	"github.com/google/uuid"
)

type Repository interface {
	Save(task *Task) error
	FindByID(id string) (*Task, error)
	FindAll() (*map[string]Task, error)
	Delete(id string) error
}

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Save(task *Task) error {
	tasks := make(map[string]Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return ErrLoadDataFailed
	}

	taskToStore := Task{
		ID:          uuid.New().String(),
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		UserId:      task.UserId,
		Status:      task.Status,
	}

	tasks[taskToStore.ID] = taskToStore

	if err := db.SaveData(db.TaskFile, tasks); err != nil {
		return ErrSaveDataFailed
	}

	return nil
}

func (r *Repo) FindByID(id string) (*Task, error) {
	tasks := make(map[string]Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return &Task{}, ErrLoadDataFailed
	}

	task, found := tasks[id]
	if !found {
		return &Task{}, ErrTaskWithGivenIdNotFound
	}

	return &task, nil
}

func (r *Repo) FindAll() (*map[string]Task, error) {
	tasks := make(map[string]Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return nil, ErrLoadDataFailed
	}

	if len(tasks) == 0 {
		return nil, ErrNoTasksFound
	}

	return &tasks, nil
}

func (r *Repo) Delete(id string) error {
	tasks := make(map[string]Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return ErrLoadDataFailed
	}

	if _, found := tasks[id]; !found {
		return ErrTaskWithGivenIdNotFound
	}

	delete(tasks, id)

	if err := db.SaveData(db.TaskFile, tasks); err != nil {
		return ErrLoadDataFailed
	}

	return nil
}
