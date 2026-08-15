package services

import (
	"ISpringTODOList/internal/appErrors"
	"ISpringTODOList/internal/models"
	"ISpringTODOList/internal/repositories"
	"strings"
)

type TaskService struct {
	repository repositories.ITaskRepository
}

func NewTaskService(repository repositories.ITaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (service *TaskService) GetActiveTasks() ([]models.Task, error) {

	return service.repository.GetActive()
}

func (service *TaskService) GetArchivedTasks() ([]models.Task, error) {

	return service.repository.GetArchived()
}

func (service *TaskService) CreateTask(task *models.Task) error {
	if strings.TrimSpace(task.Text) == "" {
		return appErrors.ErrTaskTextEmpty
	}

	// []rune, т.к len читает байты, а не символы
	if len([]rune(task.Text)) > 1000 {
		return appErrors.ErrTaskTextTooLong
	}

	return service.repository.Create(task)
}

func (service *TaskService) DeleteTask(id uint) error {

	task, err := service.repository.GetByID(id)
	if err != nil {
		return err
	}

	return service.repository.Delete(&task)
}

func (service *TaskService) CompleteTask(id uint) (models.Task, error) {
	task, err := service.repository.GetByID(id)
	if err != nil {
		return task, err
	}

	task.Completed = true
	if err := service.repository.Update(&task); err != nil {
		return task, err
	}

	return task, nil
}
