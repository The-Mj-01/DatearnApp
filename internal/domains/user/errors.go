package user

import (
	"errors"
	"fmt"
)

var (
	UserAlreadyExists  error = errors.New("a user with required phone number, already exists in database")
	UserDoesntExists   error = errors.New("requested user doesn't exists")
	WrongCredentials   error = errors.New("provided credentials are wrong")
	AuthSomethingWrong error = fmt.Errorf("%s", "Something went wrong, please contact administrator")
)
