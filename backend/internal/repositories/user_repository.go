package repository

import (
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(*models.User) (*models.User, error)
	GetOne(id int) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(*models.User) (*models.User, error)
	DeleteOne(id int) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) Create(user *models.User) (*models.User, error) {
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

func (u *UserRepository) Update(user *models.User) (*models.User, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) DeleteOne(id int) (*models.User, error) {
	var user models.User
	if err := u.db.Delete(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
