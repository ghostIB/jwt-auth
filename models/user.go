package models

import "gorm.io/gorm"

// gorm.Model автоматично додає поля: ID, CreatedAt, UpdatedAt, DeletedAt
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

// RegisterInput - дані від користувача при реєстрації
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginInput - дані при логіні
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
