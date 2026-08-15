package services

import (
	"ISpringTODOList/internal/appErrors"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/repositories"
)

type TaskService struct {
	repository repositories.ITaskRepository
}

func NewTaskService(repository repositories.ITaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (service *TaskService) GetTasks() ([]models.Task, error) {

	tasks, err := service.repository.GetAll()
	if err != nil {
		return nil, appErrors.ErrGetTasks
	}
	return tasks, nil
}

func (service *TaskService) CreateTask(task *models.Task) error {
	if task.Text == "" {
		return appErrors.ErrTaskTextEmpty
	}

	// []rune, т.к len читает байты, а не символы
	if len([]rune(task.Text)) > 1000 {
		return appErrors.ErrTaskTextTooLong
	}

	if err := service.repository.Create(task); err != nil {
		return appErrors.ErrCreateTask
	}

	return nil
}

func (service *TaskService) DeleteTask(id uint) error {

	task, err := service.repository.GetByID(id)
	if err != nil {
		return err
	}

	if err := service.repository.Delete(&task); err != nil {
		return appErrors.ErrDeleteTask
	}

	return nil
}
