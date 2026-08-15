package tests

import "ISpringTODOList/internal/models"

type mockTaskRepository struct {
}

func (m *mockTaskRepository) GetActive() ([]models.Task, error) {
	return nil, nil
}

func (m *mockTaskRepository) GetArchived() ([]models.Task, error) {
	return nil, nil
}

func (m *mockTaskRepository) GetByID(id uint) (models.Task, error) {
	return models.Task{}, nil
}

func (m *mockTaskRepository) Create(task *models.Task) error {
	return nil
}

func (m *mockTaskRepository) Delete(task *models.Task) error {
	return nil
}

func (m *mockTaskRepository) Update(task *models.Task) error {
	return nil
}
