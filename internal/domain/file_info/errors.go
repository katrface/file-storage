package file_info

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound   = errors.New("not found file info")
	ErrRepository = errors.New("file info repository error")
)

func NewNotFoundError(id uint) error {
	return fmt.Errorf("%w with id=%v", ErrNotFound, id)
}

func NewRepositoryError(err error) error {
	return fmt.Errorf("%w: %w", ErrRepository, err)
}
