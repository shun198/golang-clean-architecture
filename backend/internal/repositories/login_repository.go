package repository

import (
	"context"
	"errors"

	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errors.New("メールアドレスまたはパスワードが間違っています")
	ErrSystemError        = errors.New("システムエラーが発生しました")
)

type ILoginRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) ILoginRepository {
	return &LoginRepository{
		db: db,
	}
}

func (r *LoginRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, ErrSystemError
	}

	return &user, nil
}
