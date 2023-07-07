package image

import "errors"

var (
	ImageNotFound         error = errors.New("request image does not found")
	NameNotFound          error = errors.New("request image name does not found")
	ImageableIdNotFound   error = errors.New("request image imageableId does not found")
	ImageableTypeNotFound error = errors.New("request image imageableType does not found")
	PathNotFound          error = errors.New("request image path does not found")
)
