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
}

type locUsecase struct {
	locRepo repository.LocRepo
}

// NewLocService make and return locService
func NewLocService(locRepo repository.LocRepo) LocService {
	locUsecase := &locUsecase{locRepo}
	return locUsecase
}

func (lu *locUsecase) GetLastLocation(innerCtx *gorm.DB, id string) (*models.Location, error) {
	loc, err := lu.locRepo.GetLastLocation(innerCtx, id)

	return loc, err
}

func (lu *locUsecase) Create(innerCtx *gorm.DB, loc *models.Location) (string, error) {
	result, err := lu.locRepo.Create(innerCtx, loc)

	return result, err
}
