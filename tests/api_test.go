package tests

import (
	"ISpringTODOList/internal/database"
	"ISpringTODOList/internal/handlers"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/repositories"
	"ISpringTODOList/internal/router"
	"ISpringTODOList/internal/services"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestServer(t *testing.T) (*httptest.Server, *gorm.DB) {
	t.Helper() // маркер

	db, err := database.Connect()
	require.NoError(t, err)

	err = db.AutoMigrate(&models.Task{})
	require.NoError(t, err)

	repository := repositories.NewTaskRepository(db)
	service := services.NewTaskService(repository)
	taskHandler := handlers.NewTaskHandler(service)

	mux := router.SetupRouter(taskHandler)

	return httptest.NewServer(mux), db
}
