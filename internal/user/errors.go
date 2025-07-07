package user

import "errors"

var (
	ErrNoUsersFound   = errors.New("no users found")
	ErrUserIDRequired = errors.New("user ID is required")
	ErrLoadDataFailed = errors.New("failed to load user data")
)
