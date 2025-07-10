package task

import (
	"RestCrud/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Save(task *model.Task) error
	FindByID(id string) (*model.Task, error)
	FindAll() ([]model.Task, error)
	Delete(id string) error
}

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Save(task *model.Task) error {

	if err := r.DB.Save(task).Error; err != nil {
		return ErrSaveDataFailed
	}
	return nil
}

func (r *Repo) FindByID(id string) (*model.Task, error) {
	var task model.Task
	if err := r.DB.First(&task, "id = ?", id).Error; err != nil {
		return nil, ErrTaskWithGivenIdNotFound
	}
	return &task, nil
}

func (r *Repo) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, ErrLoadDataFailed
	}
	if len(tasks) == 0 {
		return nil, ErrNoTasksFound
	}
	return tasks, nil
}

func (r *Repo) Delete(id string) error {
	result := r.DB.Delete(&model.Task{}, "id = ?", id)
	if result.Error != nil {
		return ErrFailedToDeleteTask
	}
	if result.RowsAffected == 0 {
		return ErrTaskWithGivenIdNotFound
	}
	return nil
}
