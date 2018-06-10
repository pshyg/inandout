package controller

import (
	"inandout/api/service"
	"inandout/repository"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// InitController initialize controller
func InitController(e *echo.Echo, db *gorm.DB) {
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "inandout API is Alive!!") })

	ea := e.Group("/api")
	ev := ea.Group("/v1")
	es := ev.Group("/opo")

	locRepo := repository.NewLocrepository(db)

	ls := service.NewLocService(locRepo)
	newLocController(es, ls, db)
}

// HTTPLocHandler route location request to LocService
type HTTPLocHandler struct {
	LocService service.LocService
	Db         *gorm.DB
}

func newLocController(es *echo.Group, ls service.LocService, db *gorm.DB) {
	handler := &HTTPLocHandler{
		LocService: ls,
		Db:         db,
	}

	es.POST("/loc/:userid", handler.Create)
	es.GET("/loc/:userid", handler.GetLastLocation)
}
