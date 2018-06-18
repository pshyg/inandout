package repository

import (
	"inandout/models"

	"github.com/jinzhu/gorm"
)

// UserRepo provide methods that control database's users table
type UserRepo interface {
	Create(*gorm.DB, *models.User) (string, error)
}

// UserRepository has users table infomation
type UserRepository struct {
	conn *gorm.DB
}

// NewUserRepository make LocRepo interface
func NewUserRepository(db *gorm.DB) UserRepo {
	db = db.AutoMigrate(&models.User{})
	db.Model(&models.Location{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	return &UserRepository{db}
}

// Create register new user
func (u *UserRepository) Create(ctx *gorm.DB, user *models.User) (string, error) {
	if err := ctx.Create(user).Error; err != nil {
		return "", err
	}

	return user.ID, nil
}
