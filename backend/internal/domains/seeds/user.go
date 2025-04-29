package seed

import (
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
)

func CreateUserLocalData() []models.User {

	return []models.User{
		// システム管理者
		{
			ID:        1,
			Email:     "system1@example.com",
			Username:  "システム管理者1",
			Password:  "test",
			Role:      models.AdminRole,
			CreatedBy: 1,
			UpdatedBy: 1,
			IsActive:  true,
		},
		{
			ID:        2,
			Email:     "system2@example.com",
			Username:  "システム管理者2",
			Password:  "test",
			Role:      models.AdminRole,
			CreatedBy: 1,
			UpdatedBy: 1,
			IsActive:  true,
		},
	}
}
