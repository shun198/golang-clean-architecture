package repository

import (
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(*models.User) (*models.User, error)
	GetOne(id int) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user *models.User) (*models.User, error) {
	user.CreatedBy = 1
	user.UpdatedBy = 1

	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetOne(id int) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
