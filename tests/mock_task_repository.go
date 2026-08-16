package tests

import "ISpringTODOList/internal/models"

type mockTaskRepository struct {
	CreateCalled bool
	DeleteCalled bool
	UpdateCalled bool
	GetCalled    bool

	Tasks    []models.Task
	Task     models.Task
	GetError error
}

func (m *mockTaskRepository) GetActive() ([]models.Task, error) {
	return nil, nil
}

func (m *mockTaskRepository) GetArchived() ([]models.Task, error) {
	return nil, nil
}

func (m *mockTaskRepository) GetByID(id uint) (models.Task, error) {
	return m.Task, m.GetError
}

func (m *mockTaskRepository) Create(task *models.Task) error {
	m.CreateCalled = true
	return m.GetError
}

func (m *mockTaskRepository) Delete(task *models.Task) error {
	m.DeleteCalled = true
	return m.GetError
}

func (m *mockTaskRepository) Update(task *models.Task) error {
	m.UpdateCalled = true
	return m.GetError
}
