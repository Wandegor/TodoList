package mappers

import (
	"ISpringTODOList/internal/dto"
	"ISpringTODOList/internal/models"
)

func MapToTaskResponseDTO(task models.Task) dto.TaskResponseDTO {
	return dto.TaskResponseDTO{
		ID:        task.ID,
		Text:      task.Text,
		Completed: task.Completed,
	}
}

func MapToTaskResponseDTOs(tasks []models.Task) []dto.TaskResponseDTO {
	var response []dto.TaskResponseDTO

	for _, task := range tasks {
		response = append(response, MapToTaskResponseDTO(task))
	}

	return response
}
