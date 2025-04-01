package dto

import "todo-list/internal/domain/model"

type TaskResponseDTO struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ToTaskResponseDTO(task model.Task) TaskResponseDTO {
	return TaskResponseDTO{
		ID:      task.ID.String(),
		Title:   task.Title,
		Content: task.Content,
	}
}
