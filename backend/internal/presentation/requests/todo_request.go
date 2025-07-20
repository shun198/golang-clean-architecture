package requests

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsStarred   bool   `json:"is_starred" binding:"required"`
	IsCompleted bool   `json:"is_completed" binding:"required"`
}
