package service

import (
	"inandout/models"
	"inandout/repository"

	"github.com/jinzhu/gorm"
)

// UserService provide method for account
type UserService interface {
	Create(*gorm.DB, *models.User) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

// NewUserService make and return userService
func NewUserService(userRepo repository.UserRepo) UserService {
	userUsecase := &userUsecase{userRepo}
	return userUsecase
}

// Create register new user
func (u *userUsecase) Create(db *gorm.DB, user *models.User) (string, error) {
	innerCtx := db.New()
	result, err := u.userRepo.Create(innerCtx, user)

	return result, err
}
