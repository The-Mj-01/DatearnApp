package bio

import (
	"errors"
	"fmt"
)

var (
	BioDoesntExists    error = errors.New("requested bio doesn't exists")
	WrongCredentials   error = errors.New("provided credentials are wrong")
	AuthSomethingWrong error = fmt.Errorf("%s", "Something went wrong, please contact administrator")
)
