package repository

import "todo-list/internal/domain/model"

// Реализация нуждается в отдельной папке, infrastructure/repo/task_repo.go

// Абстракция для взаимодействия с хранилищами БД
type TaskRepository interface {
	Create(task *model.Task) error         // Создать задачу
	GetAll() ([]model.Task, error)         // Получить список всех задач
	GetByID(id string) (model.Task, error) // Получить конкретную задачу по id
	Update(task *model.Task) error         // Обновить информацию задачи (подумать, может передать id, как буду обновлять)
	Delete(id string) error                // Удалить задачу
}
