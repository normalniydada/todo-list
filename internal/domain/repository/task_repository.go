package repository

import "todo-list/internal/domain/model"

// Реализация нуждается в отдельной папке, infrastructure/repo/task_repository.go

// Абстракция для взаимодействия с хранилищами БД
type TaskRepository interface {
	Create(task *model.Task) error         // Создать задачу
	GetAll() ([]model.Task, error)         // Получить список всех задач
	GetByID(id int64) (*model.Task, error) // Получить конкретную задачу по id
	Update(task *model.Task) error         // Обновить информацию задачи (подумать, может передать id, как буду обновлять)
	Done(id int64) error
	// Выполнить задачу (по сути удаление задачи, раз она выполнена, но может просто статус изменить)
}
