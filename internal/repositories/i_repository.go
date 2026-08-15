package repositories

import "ISpringTODOList/internal/models"

type ITaskRepository interface {
	GetActive() ([]models.Task, error)
	GetArchived() ([]models.Task, error)
	GetByID(id uint) (models.Task, error)
	Create(task *models.Task) error
	Delete(task *models.Task) error
	Update(task *models.Task) error
}
