package service

import (
	"todo-list/internal/domain/model"
	"todo-list/internal/domain/repository"
)

// Реализация самих сервисов находит в domain/service (всегда), они не нуждаются в отдельной директории
// Сервисы - часть доменного слоя

// Абстракция для взаимодействия сервисов (handler)

// Я вызываю их в handlerах, при этом в хэндлерах ничего проверяю (за искл. самих handlerов)
// а проверяю в самих сервисах (которые я должен реализовать)
type TaskService interface {
	CreateTask(task *model.Task) error         // Валидация + вызов Create
	GetAllTasks() ([]model.Task, error)        // Вызов GetAll
	GetTaskByID(id int64) (*model.Task, error) // Валидация + GetByID
	UpdateTask(task *model.Task) error         // Валидация + Update
	DoneTask(id int64) error                   // Валидация + Done
}

// Примечание. Валидация - проверка входных данных

type taskServiceImpl struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskServiceImpl{repo: repo}
}

func (s *taskServiceImpl) CreateTask(task *model.Task) error {
	// s.repo.Create(task)
}

func (s *taskServiceImpl) GetAllTasks() ([]model.Task, error) {

}

func (s *taskServiceImpl) GetTaskByID(id int64) (*model.Task, error) {

}

func (s *taskServiceImpl) UpdateTask(task *model.Task) error {

}

func (s *taskServiceImpl) DoneTask(id int64) error {

}
