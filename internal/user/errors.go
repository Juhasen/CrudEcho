package user

import "errors"

var (
	ErrNoUsersFound                = errors.New("no users found")
	ErrUserIDRequired              = errors.New("user ID is required")
	ErrLoadDataFailed              = errors.New("failed to load user data")
	ErrUserIdNotFound              = errors.New("user with given ID not found")
	ErrUserAlreadyExists           = errors.New("user with given credentials already exists")
	ErrUserNameRequired            = errors.New("user name is required")
	ErrUserEmailRequired           = errors.New("user email is required")
	ErrUserEmailNotFound           = errors.New("user with given email not found")
	ErrUserIDMismatch              = errors.New("user ID mismatch")
	ErrFailedToDeleteUser          = errors.New("failed to delete user")
	ErrUserEmailInvalid            = errors.New("user email is invalid")
	ErrUserWithGivenIdDoesNotExist = errors.New("user with given ID does not exist")
	ErrAtLeastOneFieldRequired     = errors.New("at least one field (name or email) is required for update")
	ErrIdIsNotValid                = errors.New("ID is not valid UUID format")
)
