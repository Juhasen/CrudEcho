package task

import (
	"RestCrud/internal/db"
	"RestCrud/internal/task/errors"
	"RestCrud/internal/task/model"
	"github.com/google/uuid"
)

type Repository interface {
	Save(task *model.Task) error
	FindByID(id string) (*model.Task, error)
	FindAll() (*map[string]model.Task, error)
	Delete(id string) error
}

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Save(task *model.Task) error {
	tasks := make(map[string]model.Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return errors.ErrLoadDataFailed
	}

	taskToStore := model.Task{
		ID:          uuid.New().String(),
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		UserId:      task.UserId,
		Status:      task.Status,
	}

	tasks[taskToStore.ID] = taskToStore

	if err := db.SaveData(db.TaskFile, tasks); err != nil {
		return errors.ErrSaveDataFailed
	}

	return nil
}

func (r *Repo) FindByID(id string) (*model.Task, error) {
	tasks := make(map[string]model.Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return &model.Task{}, errors.ErrLoadDataFailed
	}

	task, found := tasks[id]
	if !found {
		return &model.Task{}, errors.ErrTaskWithGivenIdNotFound
	}

	return &task, nil
}

func (r *Repo) FindAll() (*map[string]model.Task, error) {
	tasks := make(map[string]model.Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return nil, errors.ErrLoadDataFailed
	}

	if len(tasks) == 0 {
		return nil, errors.ErrNoTasksFound
	}

	return &tasks, nil
}

func (r *Repo) Delete(id string) error {
	tasks := make(map[string]model.Task)
	if err := db.LoadData(db.TaskFile, &tasks); err != nil {
		return errors.ErrLoadDataFailed
	}

	if _, found := tasks[id]; !found {
		return errors.ErrTaskWithGivenIdNotFound
	}

	delete(tasks, id)

	if err := db.SaveData(db.TaskFile, tasks); err != nil {
		return errors.ErrLoadDataFailed
	}

	return nil
}
