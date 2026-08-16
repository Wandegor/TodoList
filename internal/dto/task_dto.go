package dto

type TaskRequestDTO struct {
	Text string `json:"text"`
}

type TaskResponseDTO struct {
	ID        uint   `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}
