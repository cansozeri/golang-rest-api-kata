package errors

import "errors"

var ErrNotFound = errors.New("entity not found on database")

var ErrInvalidEntity = errors.New("invalid entity")
