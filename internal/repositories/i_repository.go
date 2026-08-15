package repositories

import "ISpringTODOList/internal/models"

type ITaskRepository interface {
	GetAll() ([]models.Task, error)
	GetByID(id uint) (models.Task, error)
	Create(task *models.Task) error
	Delete(task *models.Task) error
}
