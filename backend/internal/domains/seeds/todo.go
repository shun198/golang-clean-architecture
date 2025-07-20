package seed

import (
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
)

func CreateTodoLocalData() []models.Todo {

	return []models.Todo{
		{
			ID:          1,
			Title:       "タスク1",
			Description: "タスク詳細1",
			IsStarred:   false,
			IsCompleted: false,
		},
		{
			ID:          2,
			Title:       "タスク2",
			Description: "タスク詳細2",
			IsStarred:   true,
			IsCompleted: false,
		},
		{
			ID:          3,
			Title:       "タスク3",
			Description: "タスク詳細3",
			IsStarred:   true,
			IsCompleted: true,
		},
	}
}
