package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"todo-list/internal/domain/model"
	"todo-list/internal/domain/repository"
)

var (
	ErrTaskTitleEmpty   = errors.New("task title is empty")
	ErrTaskContentEmpty = errors.New("task content is empty")
)

// Реализация самих сервисов находит в domain/service (всегда), они не нуждаются в отдельной директории
// Сервисы - часть доменного слоя

// Абстракция для взаимодействия сервисов (handler)

// Я вызываю их в handlerах, при этом в хэндлерах ничего проверяю (за искл. самих handlerов)
// а проверяю в самих сервисах (которые я должен реализовать)
type TaskService interface {
	CreateTask(ctx context.Context, title, content string) (model.Task, error)            // Валидация + вызов Create
	GetAllTasks(ctx context.Context) ([]model.Task, error)                                // Вызов GetAll
	GetTaskByID(ctx context.Context, id string) (model.Task, error)                       // Валидация + GetByID
	UpdateTask(ctx context.Context, id string, title, content string) (model.Task, error) // Валидация + Update
	DeleteTask(ctx context.Context, id string) error                                      // Валидация + Done
}

// Примечание. Валидация - проверка входных данных

type taskServiceImpl struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskServiceImpl{repo: repo}
}

func (s *taskServiceImpl) validateTaskData(title, content string) error {
	if len(title) == 0 {
		return ErrTaskTitleEmpty
	}

	if len(content) == 0 {
		return ErrTaskContentEmpty
	}

	return nil
}

func (s *taskServiceImpl) CreateTask(ctx context.Context, title, content string) (model.Task, error) {
	err := s.validateTaskData(title, content)
	if err != nil {
		return model.Task{}, err
	}

	task := model.Task{
		ID:      uuid.New(),
		Title:   title,
		Content: content,
	}
	if err = s.repo.Create(ctx, &task); err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (s *taskServiceImpl) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *taskServiceImpl) GetTaskByID(ctx context.Context, id string) (model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *taskServiceImpl) UpdateTask(ctx context.Context, id string, title, content string) (model.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return model.Task{}, err
	}

	err = s.validateTaskData(task.Title, task.Content)
	if err != nil {
		return model.Task{}, err
	}

	task.Title = title
	task.Content = content

	if err = s.repo.Update(ctx, &task); err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (s *taskServiceImpl) DeleteTask(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
