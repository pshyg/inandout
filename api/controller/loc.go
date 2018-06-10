package controller

import (
	"inandout/models"
	"net/http"

	"github.com/labstack/echo"
)

// GetLastLocation return opponent last location
func (h *HTTPLocHandler) GetLastLocation(c echo.Context) error {
	//innerCtx := h.Db.New()
	result, err := h.LocService.GetLastLocation()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, result)
}

// Update update DB loc table using request json data
func (h *HTTPLocHandler) Update(c echo.Context) error {
	innerCtx := h.Db.New()

	loc := &models.Location{UserID: c.Param("userid")}
	if err := c.Bind(loc); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.LocService.Create(innerCtx, loc)
	if err != nil {
		return c.String(http.StatusBadRequest, "InnerSystem Error")
	}

	return c.String(http.StatusOK, result)
}
