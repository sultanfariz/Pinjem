package exceptions

import "errors"

var (
	// ErrInvalidCredentials is thrown when the user credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrInternalServerError is thrown when the server encounters an error
	ErrInternalServerError = errors.New("internal server error")
	// ErrUserAlreadyExists is thrown when the user already exists
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrInsufficienBalance is thrown when the user has insufficient balance
	ErrInsufficientBalance = errors.New("insufficient balance")
	// ErrBookNotFound is thrown when the book is not found
	ErrBookNotFound = errors.New("book not found")
	// ErrUserNotFound is thrown when the user is not found
	ErrUserNotFound = errors.New("user not found")
	// ErrOrderNotFound is thrown when the order is not found
	ErrOrderNotFound = errors.New("order not found")
)
