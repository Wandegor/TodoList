package tests

import (
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/services"
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

func TestTaskService_UpdateTask(t *testing.T) {
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
	assert.Equal(t, task, result)

}
