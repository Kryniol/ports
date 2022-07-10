package domain

import "errors"

var (
	ErrInvalidArgument = errors.New("provided value is invalid")
	ErrNotFound        = errors.New("entity has not been found")
)
