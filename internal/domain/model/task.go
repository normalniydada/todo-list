package model

import "github.com/google/uuid"

type Task struct {
	ID      uuid.UUID `gorm:"primaryKey" json:"id"`   // дописать gorm
	Title   string    `gorm:"title" json:"title"`     // дописать gorm
	Content string    `gorm:"content" json:"content"` // дописать gorm
}
