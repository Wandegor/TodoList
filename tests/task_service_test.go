package tests

import (
	"ISpringTODOList/internal/appErrors"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/services"
	"errors"
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
	assert.Equal(t, task, &repo.Task)
}

func TestTaskService_CreateTask_EmptyText(t *testing.T) {
	repo := &mockTaskRepository{}
	service := services.NewTaskService(repo)

	task := &models.Task{Text: ""}

	err := service.CreateTask(task)

	assert.ErrorIs(t, err, appErrors.ErrTaskTextEmpty)
	assert.False(t, repo.CreateCalled)
}
func TestTaskService_CreateTask_OnlySpacesText(t *testing.T) {
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

func TestTaskService_CreateTask_RepositoryError(t *testing.T) {
	repoError := errors.New("repository create error")

	repo := &mockTaskRepository{
		CreateError: repoError,
	}

	service := services.NewTaskService(repo)

	task := &models.Task{
		Text: "Test task",
	}

	err := service.CreateTask(task)

	assert.ErrorIs(t, err, repoError)
	assert.True(t, repo.CreateCalled)
}

func TestTaskService_DeleteTask_RepositoryError(t *testing.T) {
	repoError := errors.New("repository delete error")

	repo := &mockTaskRepository{
		Task: models.Task{
			ID:   1,
			Text: "Test task",
		},
		DeleteError: repoError,
	}

	service := services.NewTaskService(repo)

	err := service.DeleteTask(1)

	assert.ErrorIs(t, err, repoError)
	assert.True(t, repo.DeleteCalled)
}

func TestTaskService_CompleteTask_RepositoryError(t *testing.T) {
	repoError := errors.New("repository update error")

	repo := &mockTaskRepository{
		Task: models.Task{
			ID:        1,
			Text:      "Test task",
			Completed: false,
		},
		UpdateError: repoError,
	}

	service := services.NewTaskService(repo)

	result, err := service.CompleteTask(1)

	assert.ErrorIs(t, err, repoError)
	assert.Nil(t, result)
	assert.True(t, repo.UpdateCalled)
}
