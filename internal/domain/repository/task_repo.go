package repository

import (
	"context"
	"todo-list/internal/domain/model"
)

// Реализация нуждается в отдельной папке, infrastructure/repo/task_repo.go

// Абстракция для взаимодействия с хранилищами БД
type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	GetAll(ctx context.Context) ([]model.Task, error)
	GetByID(ctx context.Context, id string) (model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id string) error
}
