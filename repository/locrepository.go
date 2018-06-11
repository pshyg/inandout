package repository

import (
	"fmt"
	"inandout/models"

	"github.com/jinzhu/gorm"
)

// LocRepo provide method that controll dataBase
type LocRepo interface {
	Create(*gorm.DB, *models.Location) (string, error)
	GetLastLocation(*gorm.DB, string) (*models.Location, error)
	GetAllOpponentLoc(*gorm.DB, string) (*models.Locations, error)
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
	var count int
	innerCtx.Model(&models.Location{}).Where("time = ? AND user_id = ?", loc.Time, loc.UserID).Count(&count)

	if count > 0 {
		return "", fmt.Errorf("duplicated time at same userID")
	}

	if err := innerCtx.Create(loc).Error; err != nil {
		return "", err
	}

	return loc.UserID, nil
}

// GetLastLocation return opponent's last location
func (l LocRepository) GetLastLocation(innerCtx *gorm.DB, id string) (*models.Location, error) {
	loc := &models.Location{}
	scope := innerCtx.Model(&models.Location{}).Limit(1).Where("user_id NOT LIKE ?", id).Order("time desc").Find(&loc)

	if scope.Error != nil {
		return nil, scope.Error
	} else if scope.RowsAffected == 0 {
		return nil, fmt.Errorf("record not found")
	}

	return loc, nil
}

// GetAllOpponentLoc return all data of oppoent
func (l LocRepository) GetAllOpponentLoc(innerCtx *gorm.DB, id string) (*models.Locations, error) {
	var locSlice []*models.Location

	scope := innerCtx.Model(&models.Location{}).Where("user_id NOT LIKE ?", id).Order("time desc").Find(&locSlice)

	if scope.Error != nil {
		return nil, scope.Error
	} else if scope.RowsAffected == 0 {
		return nil, fmt.Errorf("record not found")
	}

	locs := &models.Locations{locSlice}

	return locs, nil
}
