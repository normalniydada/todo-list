package model

import "github.com/google/uuid"

type Task struct {
	ID      uuid.UUID `gorm:"primaryKey" json:"id"`
	Title   string    `gorm:"title" json:"title"`
	Content string    `gorm:"content" json:"content"`
}
