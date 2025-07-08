package user

import (
	"RestCrud/internal/user/dto"
	"RestCrud/internal/user/model"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user *dto.UserResponseDTO) error
	FindByID(id string) (*model.User, error)
	FindAll() ([]model.User, error)
	Delete(id string) error
	FindByEmail(email string) (*model.User, error)
}

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}

func (r *Repo) Save(user *dto.UserResponseDTO) error {
	userToStore := model.User{
		Name:  user.Name,
		Email: user.Email,
	}

	if err := r.DB.Create(&userToStore).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repo) FindByID(id string) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, ErrNoUsersFound
	}
	return users, nil
}

func (r *Repo) Delete(id string) error {
	if err := r.DB.Delete(&model.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, ErrUserEmailNotFound
	}

	return &user, nil
}
