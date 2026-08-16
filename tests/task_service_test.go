package tests

import (
	"ISpringTODOList/internal/appErrors"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/services"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskService_CreateTask(t *testing.T) {
	repo := &mockTaskRepository{}
	service := services.NewTaskService(repo)

	task := &models.Task{Text: "TestTaskService_CreateTask"}

	err := service.CreateTask(task)

	assert.Nil(t, err)
	assert.True(t, repo.CreateCalled)
}

func TestTaskService_CreateTask_EmptyText(t *testing.T) {
	repo := &mockTaskRepository{}
	service := services.NewTaskService(repo)

	task := &models.Task{Text: "    "}

	err := service.CreateTask(task)

	assert.ErrorIs(t, err, appErrors.ErrTaskTextEmpty)
	assert.False(t, repo.CreateCalled)
}

func TestTaskService_CreateTask_TooLongText(t *testing.T) {
	repo := &mockTaskRepository{}
	service := services.NewTaskService(repo)

	task := &models.Task{Text: strings.Repeat("r", 1001)}

	err := service.CreateTask(task)

	assert.ErrorIs(t, err, appErrors.ErrTaskTextTooLong)
	assert.False(t, repo.CreateCalled)
}

func TestTaskService_DeleteTask(t *testing.T) {
	repo := &mockTaskRepository{
		Task: models.Task{
			ID:        1,
			Text:      "To_Delete",
			Completed: false,
		},
	}

	service := services.NewTaskService(repo)

	err := service.DeleteTask(1)

	assert.Nil(t, err)
	assert.True(t, repo.DeleteCalled)
}

func TestTaskService_DeleteTask_NotFound(t *testing.T) {
	repo := &mockTaskRepository{
		GetError: appErrors.ErrTaskNotFound,
	}

	service := services.NewTaskService(repo)

	err := service.DeleteTask(900)

	assert.ErrorIs(t, err, appErrors.ErrTaskNotFound)
	assert.False(t, repo.DeleteCalled)
}

func TestTaskService_CompleteTask(t *testing.T) {
	task := models.Task{
		ID:        1,
		Text:      "To_Complete",
		Completed: false,
	}
	repo := &mockTaskRepository{Task: task}

	service := services.NewTaskService(repo)

	task.Completed = true

	result, err := service.CompleteTask(1)

	assert.Nil(t, err)
	assert.True(t, repo.UpdateCalled)
	assert.Equal(t, &task, result)
}

func TestTaskService_CompleteTask_NotFound(t *testing.T) {
	repo := &mockTaskRepository{
		GetError: appErrors.ErrTaskNotFound,
	}

	service := services.NewTaskService(repo)

	result, err := service.CompleteTask(900)

	assert.ErrorIs(t, err, appErrors.ErrTaskNotFound)
	assert.False(t, repo.UpdateCalled)
	assert.Nil(t, result)
}
