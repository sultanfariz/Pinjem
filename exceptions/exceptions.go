package exceptions

import "errors"

var (
	// ErrInvalidCredentials is thrown when the user credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrInternalServerError is thrown when the server encounters an error
	ErrInternalServerError = errors.New("internal server error")
	// ErrUserAlreadyExists is thrown when the user already exists
	ErrUserAlreadyExists = errors.New("user already exists")
)
