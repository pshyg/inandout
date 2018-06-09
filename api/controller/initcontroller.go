package controller

import (
	"inandout/api/service"
	"net/http"

	"github.com/labstack/echo"
)

// InitController initialize controller
func InitController(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "inandout API is Alive!!") })

	ea := e.Group("/api")
	ev := ea.Group("/v1")
	es := ev.Group("/opo")

	ls := service.NewLocService()
	newLocController(es, ls)
}

// HTTPLocHandler route location request to LocService
type HTTPLocHandler struct {
	LocService service.LocService
}

func newLocController(es *echo.Group, ls service.LocService) {
	handler := &HTTPLocHandler{
		LocService: ls,
	}

	es.GET("/loc", handler.GetLastLocation)
}
