package service

import (
	"inandout/models"
	"inandout/repository"

	"github.com/jinzhu/gorm"
)

// LocService provide method for location
type LocService interface {
	GetLastLocation(*gorm.DB, string) (*models.Location, error)
	Create(*gorm.DB, *models.Location) (string, error)
	GetAllOpponentLoc(*gorm.DB, string) ([]*models.Location, error)
}

type locUsecase struct {
	locRepo repository.LocRepo
}

// NewLocService make and return locService
func NewLocService(locRepo repository.LocRepo) LocService {
	locUsecase := &locUsecase{locRepo}
	return locUsecase
}

// GetLastLocation return opponent's last location
func (lu *locUsecase) GetLastLocation(db *gorm.DB, id string) (*models.Location, error) {
	innerCtx := db.New()
	defer innerCtx.Close()

	loc, err := lu.locRepo.GetLastLocation(innerCtx, id)

	return loc, err
}

// Create create new location data record
func (lu *locUsecase) Create(db *gorm.DB, loc *models.Location) (string, error) {
	innerCtx := db.New()
	defer innerCtx.Close()

	result, err := lu.locRepo.Create(innerCtx, loc)

	return result, err
}

// GetAllOpponentLoc return All data of opponent
func (lu *locUsecase) GetAllOpponentLoc(db *gorm.DB, id string) ([]*models.Location, error) {
	innerCtx := db.New()
	defer innerCtx.Close()

	locs, err := lu.locRepo.GetAllOpponentLoc(innerCtx, id)

	return locs, err
}
