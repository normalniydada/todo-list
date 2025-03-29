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

func (t *taskRepositoryImpl) Create(task *model.Task) error {

}

func (t *taskRepositoryImpl) GetAll() ([]model.Task, error) {

}

func (t *taskRepositoryImpl) GetByID(id int64) (*model.Task, error) {

}

func (t *taskRepositoryImpl) Update(task *model.Task) error {

}

func (t *taskRepositoryImpl) Done(id int64) error {

}
