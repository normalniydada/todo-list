package model

type Task struct {
	ID      uint   `json:"id"`      // дописать gorm
	Title   string `json:"title"`   // дописать gorm
	Content string `json:"content"` // дописать gorm
	Done    bool   `json:"done"`    // дописать gorm
}
