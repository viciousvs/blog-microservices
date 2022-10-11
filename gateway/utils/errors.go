package utils

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrEmptyUUID     = errors.New("empty uuid")
	ErrEmptyBody     = errors.New("empty body")
	ErrInvalidUUID   = errors.New("invalid uuid")
	ErrNotValid      = errors.New("json not valid")
	ErrDoesntCreated = errors.New("post doesnt created")
	ErrDoesntUpdated = errors.New("post doesnt updated")
)
