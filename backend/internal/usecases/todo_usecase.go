package usecase

import (
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"github.com/shun198/golang-clean-architecture/internal/presentation/requests"
	repository "github.com/shun198/golang-clean-architecture/internal/repositories"
)

type ITodoUsecase interface {
	CreateTodo(req requests.CreateTodoRequest) (*models.Todo, error)
	GetTodo(id int) (*models.Todo, error)
	GetAllTodos() (*[]models.Todo, error)
	UpdateTodo(req requests.UpdateTodoRequest, todo *models.Todo) (*models.Todo, error)
	DeleteTodo(id int) (*models.Todo, error)
}

type TodoUsecase struct {
	todoRepository repository.ITodoRepository
}

func NewTodoUsecase(todoRepository repository.ITodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepository: todoRepository,
	}
}

func (u *TodoUsecase) GetAllTodos() (*[]models.Todo, error) {
	return u.todoRepository.GetAll()
}

func (u *TodoUsecase) CreateTodo(req requests.CreateTodoRequest) (*models.Todo, error) {
	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
	}
	return u.todoRepository.Create(todo)
}

func (u *TodoUsecase) GetTodo(id int) (*models.Todo, error) {
	return u.todoRepository.GetOne(id)
}

func (u *TodoUsecase) UpdateTodo(req requests.UpdateTodoRequest, todo *models.Todo) (*models.Todo, error) {
	todo.Title = req.Title
	todo.Description = req.Description
	todo.IsStarred = req.IsStarred
	todo.IsCompleted = req.IsCompleted
	return u.todoRepository.Update(todo)
}

func (u *TodoUsecase) DeleteTodo(id int) (*models.Todo, error) {
	return u.todoRepository.DeleteOne(id)
}
