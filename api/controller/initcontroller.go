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

	locRepo := repository.NewLocRepository(db)
	userRepo := repository.NewUserRepository(db)

	ls := service.NewLocService(locRepo)
	us := service.NewUserService(userRepo)

	newLocController(es, ls, db)
	newUserController(es, us, db)
}

// HTTPLocHandler route location request to LocService
type HTTPLocHandler struct {
	LocService service.LocService
	DB         *gorm.DB
}

func newLocController(es *echo.Group, ls service.LocService, db *gorm.DB) {
	handler := &HTTPLocHandler{
		LocService: ls,
		DB:         db,
	}

	es.POST("/loc/:userid", handler.Create)
	es.GET("/loc/:userid", handler.GetLastLocation)
	es.GET("/locs/:userid", handler.GetAllOpponentLoc)
}

// HTTPUserHandler route account request to UserService
type HTTPUserHandler struct {
	UserService service.UserService
	DB          *gorm.DB
}

func newUserController(es *echo.Group, us service.UserService, db *gorm.DB) {
	handler := &HTTPUserHandler{
		UserService: us,
		DB:          db,
	}

	es.POST("/user/:userid", handler.Register)
}
