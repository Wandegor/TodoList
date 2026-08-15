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
}
