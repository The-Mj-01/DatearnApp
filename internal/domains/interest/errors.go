package interest

import "errors"

var (
	InterestNotFound error = errors.New("request interest does not found")
	NameNotFound     error = errors.New("request interest name does not found")
)
