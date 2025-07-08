package user

import (
	"RestCrud/internal/user/dto"
	"RestCrud/internal/user/model"
	"github.com/google/uuid"
)

type Repository interface {
	Save(user *dto.UserDTO) error
	FindByID(id string) (*model.User, error)
	FindAll() (*map[string]model.User, error)
	Delete(id string) error
	FindByEmail(email string) (*model.User, error)
	Update(id string, user *dto.UserUpdateDTO) error
}

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Save(user *dto.UserDTO) error {
	users := make(map[string]model.User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return ErrLoadDataFailed
	}

	userToStore := model.User{
		ID:    uuid.New().String(),
		Name:  user.Name,
		Email: user.Email,
	}

	users[userToStore.ID] = userToStore

	if err := db.SaveData(db.UserFile, users); err != nil {
		return err
	}

	return nil
}

func (r *Repo) FindByID(id string) (*model.User, error) {
	users := make(map[string]model.User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return nil, ErrLoadDataFailed
	}

	user, found := users[id]
	if !found {
		return nil, ErrUserIdNotFound
	}

	return &user, nil
}

func (r *Repo) FindAll() (*map[string]model.User, error) {
	users := make(map[string]model.User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return nil, ErrLoadDataFailed
	}

	if len(users) == 0 {
		return nil, ErrNoUsersFound
	}

	return &users, nil
}

func (r *Repo) Delete(id string) error {
	users := make(map[string]model.User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return err
	}

	if _, found := users[id]; !found {
		return ErrUserIdNotFound
	}

	delete(users, id)

	if err := db.SaveData(db.UserFile, users); err != nil {
		return err
	}

	return nil
}

func (r *Repo) Update(id string, userRequest *dto.UserUpdateDTO) error {
	users := make(map[string]model.User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return ErrLoadDataFailed
	}

	user, err := r.FindByID(id)
	if err != nil {
		return err
	}

	if userRequest.Name != nil {
		user.Name = *userRequest.Name
	}

	if userRequest.Email != nil {
		user.Email = *userRequest.Email
	}

	users[id] = *user

	if err := db.SaveData(db.UserFile, users); err != nil {
		return err
	}

	return nil
}

func (r *Repo) FindByEmail(email string) (*model.User, error) {
	users := make(map[string]model.User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return nil, ErrLoadDataFailed
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, ErrUserEmailNotFound
}
