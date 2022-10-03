package utils

import "errors"

var (
	ErrNilMutex  = errors.New("Nil mutex error")
	ErrNilDB     = errors.New("Nil DB error")
	ErrEmptyUUID = errors.New("Empty uuid")
)
