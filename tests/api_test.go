package tests

import (
	"ISpringTODOList/internal/config"
	"ISpringTODOList/internal/database"
	"ISpringTODOList/internal/dto"
	"ISpringTODOList/internal/handlers"
	"ISpringTODOList/internal/handlers/router"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/repositories"
	"ISpringTODOList/internal/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestServer(t *testing.T) (*httptest.Server, *gorm.DB) {
	t.Helper() // маркер

	// TODO: исправить хардкод, (сделать бд для тестов?)
	testDBConfig := config.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "wander",
		Password: "wandersPassword",
		Name:     "todoList",
	}
	db, err := database.Connect(testDBConfig)
	require.NoError(t, err)

	err = db.AutoMigrate(&models.Task{})
	require.NoError(t, err)

	repository := repositories.NewTaskRepository(db)
	service := services.NewTaskService(repository)
	taskHandler := handlers.NewTaskHandler(service)

	mux := router.SetupRouter(taskHandler)

	return httptest.NewServer(mux), db
}

func cleanupTasks(t *testing.T, db *gorm.DB, ids []uint) {
	t.Helper()

	for _, id := range ids {
		err := db.Delete(&models.Task{}, id).Error
		require.NoError(t, err)
	}
}

func TestAPI_CreateTask(t *testing.T) {
	server, db := setupTestServer(t)
	defer server.Close()

	var createdIDs []uint
	defer func() {
		cleanupTasks(t, db, createdIDs)
	}()

	request := dto.TaskRequestDTO{
		Text: "API CreateTask test",
	}

	jsonBody, err := json.Marshal(request)
	require.NoError(t, err)

	var task dto.TaskResponseDTO

	resp, err := http.Post(
		server.URL+"/tasks",
		"application/json",
		bytes.NewReader(jsonBody),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	err = json.NewDecoder(resp.Body).Decode(&task)
	require.NoError(t, err)

	assert.NotZero(t, task.ID)
	assert.Equal(t, request.Text, task.Text)
	assert.False(t, task.Completed)

	createdIDs = append(createdIDs, task.ID)
}

func TestAPI_CreateTask_InvalidJson(t *testing.T) {
	server, db := setupTestServer(t)
	defer server.Close()

	var createdIDs []uint
	defer func() {
		cleanupTasks(t, db, createdIDs)
	}()

	body := `"{ text":"API test task"`

	resp, err := http.Post(
		server.URL+"/tasks",
		"application/json",
		bytes.NewBufferString(body),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAPI_CreateTask_EmptyText(t *testing.T) {
	server, db := setupTestServer(t)
	defer server.Close()

	var createdIDs []uint
	defer func() {
		cleanupTasks(t, db, createdIDs)
	}()

	request := dto.TaskRequestDTO{
		Text: "",
	}

	jsonBody, err := json.Marshal(request)
	require.NoError(t, err)

	resp, err := http.Post(
		server.URL+"/tasks",
		"application/json",
		bytes.NewReader(jsonBody),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAPI_CreateTask_WhitespaceText(t *testing.T) {
	server, db := setupTestServer(t)
	defer server.Close()

	var createdIDs []uint
	defer func() {
		cleanupTasks(t, db, createdIDs)
	}()

	request := dto.TaskRequestDTO{
		Text: "     ",
	}

	jsonBody, err := json.Marshal(request)
	require.NoError(t, err)

	resp, err := http.Post(
		server.URL+"/tasks",
		"application/json",
		bytes.NewReader(jsonBody),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAPI_CreateTask_TooLongText(t *testing.T) {
	server, db := setupTestServer(t)
	defer server.Close()

	var createdIDs []uint
	defer func() {
		cleanupTasks(t, db, createdIDs)
	}()

	request := dto.TaskRequestDTO{
		Text: strings.Repeat("r", 1001),
	}

	jsonBody, err := json.Marshal(request)
	require.NoError(t, err)

	resp, err := http.Post(
		server.URL+"/tasks",
		"application/json",
		bytes.NewReader(jsonBody),
	)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
