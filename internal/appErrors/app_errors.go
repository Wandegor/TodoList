package appErrors

import "errors"

var (
	ErrInternalServerError = errors.New("internal Server Error")
	ErrInvalidRequestBody  = errors.New("invalid request body")
	ErrTaskNotFound        = errors.New("task not found")
	ErrTaskTextEmpty       = errors.New("task text is empty")
	ErrTaskTextTooLong     = errors.New("task text is too long")
	ErrInvalidTaskID       = errors.New("invalid task id")

	ErrCreateTask = errors.New("failed to create task")
	ErrGetTasks   = errors.New("failed to get tasks")
	ErrDeleteTask = errors.New("failed to delete tasks")
)
