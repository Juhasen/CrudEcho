package user

import "errors"

var (
	ErrNoUsersFound      = errors.New("no users found")
	ErrUserIDRequired    = errors.New("user ID is required")
	ErrLoadDataFailed    = errors.New("failed to load user data")
	ErrUserIdNotFound    = errors.New("user with given ID not found")
	ErrUserAlreadyExists = errors.New("user with given ID already exists")
	ErrUserNameRequired  = errors.New("user name is required")
	ErrUserEmailRequired = errors.New("user email is required")
)
