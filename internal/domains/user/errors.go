package user

import "errors"

var (
	UserAlreadyExists error = errors.New("a user with required phone number, already exists in database")
	UserDoesntExists  error = errors.New("requested user doesn't exists")
)
