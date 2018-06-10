package repository

import (
	"fmt"
	"inandout/models"

	"github.com/jinzhu/gorm"
)

// LocRepo provide method that controll dataBase
type LocRepo interface {
	Create(*gorm.DB, *models.Location) (string, error)
}

// LocRepository has location table infomation
type LocRepository struct {
	conn *gorm.DB
}

// NewLocrepository make LocRepo interface
func NewLocrepository(db *gorm.DB) LocRepo {
	db = db.AutoMigrate(&models.Location{})

	return &LocRepository{db}
}

// Create create new record database Location table
func (l LocRepository) Create(innerCtx *gorm.DB, loc *models.Location) (string, error) {
	if !innerCtx.NewRecord(*loc) {
		return "", fmt.Errorf("duplicated record")
	}

	if err := innerCtx.Create(loc).Error; err != nil {
		return "", err
	}

	return loc.UserID, nil
}
