package service

import "errors"

var (
	ErrNotFound    = errors.New("not found")
	ErrNotFoundCmd = errors.New("not found cmd")
)
