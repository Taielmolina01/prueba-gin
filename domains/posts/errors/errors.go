package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type ErrorPostAlreadyExist struct {
	ID uuid.UUID
}

func (e ErrorPostAlreadyExist) Error() string {
	return fmt.Sprintf("Post with id %d already exists", e.ID)
}

type ErrorPostDoesntExist struct {
	ID uuid.UUID
}

func (e ErrorPostDoesntExist) Error() string {
	return fmt.Sprintf("Post with id %d not exists", e.ID)
}

type ErrorPostWithoutContent struct{}

func (e ErrorPostWithoutContent) Error() string {
	return "Posts must have content"
}

type ErrorPostWithoutTitle struct{}

func (e ErrorPostWithoutTitle) Error() string {
	return "Posts must have title"
}
