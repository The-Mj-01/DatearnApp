package location

import (
	"errors"
)

var (
	CountryNotFound error = errors.New("request country does not found")
	CityNotFound    error = errors.New("request city does not found")

	//UserIdNotFound     error = errors.New("request userId does not found")
	//BornNotFound       error = errors.New("request born does not found")
	//YouAreNotAllowed   error = errors.New("operation is not allowed for you")
	//WrongCredentials   error = errors.New("provided credentials are wrong")
	//AuthSomethingWrong error = fmt.Errorf("%s", "Something went wrong, please contact administrator")
)
