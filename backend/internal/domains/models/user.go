package models

import (
	"time"
)

const (
	AdminRole   = "admin"
	GeneralRole = "general"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email" validate:"required,email"`
	Username  string    `gorm:"not null" json:"username" validate:"required"`
	Password  string    `gorm:"not null" json:"password" validate:"required"`
	Role      string    `gorm:"not null" json:"role" validate:"required,oneof=admin general"`
	CreatedBy int       `gorm:"not null" json:"created_by" validate:"required"`
	UpdatedBy int       `gorm:"not null" json:"updated_by" validate:"required"`
	IsActive  bool      `gorm:"default:true;not null" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ListUsersResult struct {
	Users []User
	Total int64
}
