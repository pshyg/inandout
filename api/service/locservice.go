package service

import (
	"inandout/models"
	"inandout/repository"

	"github.com/jinzhu/gorm"
)

// LocService provide method for location
type LocService interface {
	GetLastLocation() (string, error) // (*models.Location) (*models.Location, error)
	Create(*gorm.DB, *models.Location) (string, error)
}

type locUsecase struct {
	locRepo repository.LocRepo
}

// NewLocService make and return locService
func NewLocService(locRepo repository.LocRepo) LocService {
	locUsecase := &locUsecase{locRepo}
	return locUsecase
}

func (lu *locUsecase) GetLastLocation() (string, error) { // (loc *models.Location) (*models.Location, error) {
	return "123.456789, 43.2198765", nil
}

func (lu *locUsecase) Create(innerCtx *gorm.DB, loc *models.Location) (string, error) {
	result, err := lu.locRepo.Create(innerCtx, loc)

	return result, err
}
