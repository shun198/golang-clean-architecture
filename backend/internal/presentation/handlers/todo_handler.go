package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shun198/golang-clean-architecture/internal/presentation/requests"
	"github.com/shun198/golang-clean-architecture/internal/presentation/responses"
	usecase "github.com/shun198/golang-clean-architecture/internal/usecases"
)

type TodoHandler struct {
	todoUsecase usecase.ITodoUsecase
}

func NewTodoHandler(todoUsecase usecase.ITodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	results, err := h.todoUsecase.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var todos []responses.TodoResponse
	for _, todo := range results {
		todos = append(todos, responses.TodoResponse{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			IsStarred:   todo.IsStarred,
			IsCompleted: todo.IsCompleted,
		})
	}

	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req requests.CreateTodoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	todo, err := h.todoUsecase.CreateTodo(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responses.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		IsStarred:   todo.IsStarred,
		IsCompleted: todo.IsCompleted,
	})
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not a valid integer",
		})
		return
	}
	todo, err := h.todoUsecase.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "todo not found",
		})
		return
	}
	c.JSON(http.StatusOK, responses.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		IsStarred:   todo.IsStarred,
		IsCompleted: todo.IsCompleted,
	})
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var req requests.UpdateTodoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not a valid integer",
		})
		return
	}
	todo, err := h.todoUsecase.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "todo not found",
		})
		return
	}
	updated_todo, err := h.todoUsecase.UpdateTodo(req, todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, responses.TodoResponse{
		ID:          updated_todo.ID,
		Title:       updated_todo.Title,
		Description: updated_todo.Description,
		IsStarred:   updated_todo.IsStarred,
		IsCompleted: updated_todo.IsCompleted,
	})
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not a valid integer",
		})
		return
	}
	_, err = h.todoUsecase.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "todo not found",
		})
		return
	}
	_, err = h.todoUsecase.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}
