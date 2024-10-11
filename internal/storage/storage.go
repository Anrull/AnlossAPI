package storage

import "errors"

var (
	ErrNotFound  = errors.New("not found")
	ErrUrlExists = errors.New("url exists")
)
