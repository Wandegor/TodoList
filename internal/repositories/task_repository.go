package repositories

import (
	"ISpringTODOList/internal/appErrors"
	"ISpringTODOList/internal/models"
	"errors"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (repo *taskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	err := repo.db.Find(&tasks).Error

	return tasks, err
}

func (repo *taskRepository) GetByID(id uint) (models.Task, error) {
	var task models.Task

	err := repo.db.First(&task, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return task, appErrors.ErrTaskNotFound
	}

	return task, err
}

func (repo *taskRepository) Create(task *models.Task) error {
	return repo.db.Create(task).Error
}

func (repo *taskRepository) Delete(task *models.Task) error {
	return repo.db.Delete(task).Error
}
