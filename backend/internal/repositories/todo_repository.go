package repository

import (
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"gorm.io/gorm"
)

type ITodoRepository interface {
	Create(*models.Todo) (*models.Todo, error)
	GetOne(id int) (*models.Todo, error)
	GetAll() ([]models.Todo, error)
	Update(*models.Todo) (*models.Todo, error)
	DeleteOne(id int) (*models.Todo, error)
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) Create(todo *models.Todo) (*models.Todo, error) {
	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) GetOne(id int) (*models.Todo, error) {
	var todo models.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Update(todo *models.Todo) (*models.Todo, error) {
	if err := r.db.Save(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) DeleteOne(id int) (*models.Todo, error) {
	var todo models.Todo
	if err := r.db.Delete(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
