package repository

import (
	"errors"

	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByID(id int) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("ユーザが見つかりません")
		}
		return nil, err
	}
	return &user, nil
}
