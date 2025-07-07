package user

import (
	"RestCrud/internal/db"
	"errors"
	"fmt"
)

type Repository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
	FindAll() (*map[string]User, error)
	Delete(id string) error
}

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Save(user *User) error {
	users := make(map[string]User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return err
	}

	if user.ID == "" {
		return errors.New("error: user ID cannot be empty")
	}

	if _, exists := users[user.ID]; exists {
		return fmt.Errorf("error: user with ID:%s already exists", user.ID)
	}

	users[user.ID] = *user

	if err := db.SaveData(db.UserFile, users); err != nil {
		return err
	}

	return nil
}

func (r *Repo) FindByID(id string) (*User, error) {
	users := make(map[string]User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return &User{}, err
	}

	user, found := users[id]
	if !found {
		return &User{}, fmt.Errorf("error: user with ID:%s not found", id)
	}

	return &user, nil
}

func (r *Repo) FindAll() (*map[string]User, error) {
	users := make(map[string]User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("error: no users found")
	}

	return &users, nil
}

func (r *Repo) Delete(id string) error {
	users := make(map[string]User)
	if err := db.LoadData(db.UserFile, &users); err != nil {
		return err
	}

	if _, found := users[id]; !found {
		return fmt.Errorf("error: user with ID:%s not found", id)
	}

	delete(users, id)

	if err := db.SaveData(db.UserFile, users); err != nil {
		return err
	}

	return nil
}
