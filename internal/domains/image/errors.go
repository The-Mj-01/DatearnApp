package image

import "errors"

var (
	ImageNotFound error = errors.New("request image does not found")
)
