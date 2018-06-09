package service

// LocService provide method for location
type LocService interface {
	GetLastLocation() (string, error) // (*models.Location) (*models.Location, error)
}

type locUsecase struct{}

// NewLocService make and return locService
func NewLocService() LocService {
	locUsecase := new(locUsecase)
	return locUsecase
}

func (lu *locUsecase) GetLastLocation() (string, error) { // (loc *models.Location) (*models.Location, error) {
	return "123.456789, 43.2198765", nil
}
