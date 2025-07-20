package responses

type TodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsStarred   bool   `json:"is_starred"`
	IsCompleted bool   `json:"is_completed"`
}

type ListTodoResponse struct {
	Count  int64          `json:"count"`
	Length int            `json:"length"`
	Todos  []TodoResponse `json:"todos"`
}
