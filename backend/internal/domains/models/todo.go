package models

type Todo struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title" validate:"required"`
	Description string `gorm:"not null" json:"description" validate:"required"`
	IsStarred   bool   `gorm:"default:false;not null" json:"is_starred"`
	IsCompleted bool   `gorm:"default:false;not null" json:"is_completed"`
}

type ListTodosResult struct {
	Todos []Todo
	Total int64
}
