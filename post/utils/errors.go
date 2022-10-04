package utils

import "errors"

var (
	ErrNilDB     = errors.New("Nil DB error")
	ErrEmptyUUID = errors.New("Empty uuid")
	ErrNotExist  = errors.New("Element not exist")
)
