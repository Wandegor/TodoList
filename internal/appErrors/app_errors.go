package appErrors

import "errors"

var (
	ErrMethodNotAllowed    = errors.New("method not allowed")
	ErrInternalServerError = errors.New("internal Server Error")
	ErrInvalidRequestBody  = errors.New("invalid request body")
	ErrTaskNotFound        = errors.New("task not found")
	ErrTaskTextEmpty       = errors.New("task text is empty")
	ErrTaskTextTooLong     = errors.New("task text is too long")
	ErrInvalidTaskID       = errors.New("invalid task id")

	ErrGetTasks = errors.New("failed to get tasks")
)
