package repository

import (
	"fmt"
	"inandout/models"

	"github.com/jinzhu/gorm"
)

// LocRepo provide methods that control database's locations table
type LocRepo interface {
	Create(*gorm.DB, *models.Location) (string, error)
	GetLastLocation(*gorm.DB, string) (*models.Location, error)
	GetAllOpponentLoc(*gorm.DB, string) (*models.Locations, error)
}

// LocRepository has locations table infomation
type LocRepository struct {
	conn *gorm.DB
}

// NewLocRepository make LocRepo interface
func NewLocRepository(db *gorm.DB) LocRepo {
	db = db.AutoMigrate(&models.Location{})

	return &LocRepository{db}
}

// Create create new record database Location table
func (l LocRepository) Create(ctx *gorm.DB, loc *models.Location) (string, error) {
	if err := ctx.Create(loc).Error; err != nil {
		return "", err
	}

	return loc.UserID, nil
}

// GetLastLocation return opponent's last location
func (l LocRepository) GetLastLocation(ctx *gorm.DB, id string) (*models.Location, error) {
	loc := &models.Location{}
	scope := ctx.Model(&models.Location{}).Limit(1).Where("user_id NOT LIKE ?", id).Order("time desc").Find(&loc)

	if scope.Error != nil {
		return nil, scope.Error
	} else if scope.RowsAffected == 0 {
		return nil, fmt.Errorf("record not found")
	}

	return loc, nil
}

// GetAllOpponentLoc return all data of oppoent
func (l LocRepository) GetAllOpponentLoc(ctx *gorm.DB, id string) (*models.Locations, error) {
	locs := &models.Locations{}

	scope := ctx.Model(&models.Location{}).Where("user_id NOT LIKE ?", id).Order("time desc").Find(&locs.Locs)

	if scope.Error != nil {
		return nil, scope.Error
	} else if scope.RowsAffected == 0 {
		return nil, fmt.Errorf("record not found")
	}

	return locs, nil
}
