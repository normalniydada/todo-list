package repository

import (
	"gorm.io/gorm"
	"todo-list/internal/domain/model"
	"todo-list/internal/domain/repository"
)

// Вся логика взаимодействия с базой данных

type taskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &taskRepositoryImpl{db: db}
}

func (r *taskRepositoryImpl) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepositoryImpl) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepositoryImpl) GetByID(id string) (model.Task, error) {
	var task model.Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepositoryImpl) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&model.Task{}, "id = ?", id).Error
}
