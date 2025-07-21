package repository

import (
	"github.com/shun198/golang-clean-architecture/internal/domains/models"
	"github.com/shun198/golang-clean-architecture/internal/presentation/requests"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(*models.User) (*models.User, error)
	GetOne(id int) (*models.User, error)
	GetAll(params requests.ListUsersQuery) (*models.ListUsersResult, error)
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

func (r *UserRepository) GetAll(params requests.ListUsersQuery) (*models.ListUsersResult, error) {
	var users []models.User
	var total int64
	query := r.db.Model(&models.User{})
	query = applyUserFilters(query, params)
	query = query.Order("id ASC")
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	if err := query.Offset(params.Offset).Limit(params.Limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return &models.ListUsersResult{
		Users: users,
		Total: total,
	}, nil
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetOne(id int) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteOne(id int) (*models.User, error) {
	var user models.User
	if err := r.db.Delete(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// sqlインジェクション対策(https://gorm.io/ja_JP/docs/security.html)
func applyUserFilters(db *gorm.DB, params requests.ListUsersQuery) *gorm.DB {
	if params.Username != "" {
		db = db.Where("username LIKE ?", "%"+params.Username+"%")
	}
	if params.Email != "" {
		db = db.Where("email LIKE ?", "%"+params.Email+"%")
	}
	if params.Role != "" {
		db = db.Where("role = ?", "%"+params.Role+"%")
	}
	return db
}
