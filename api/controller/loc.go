package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetLastLocation return opponent last location
func (h *HTTPLocHandler) GetLastLocation(c echo.Context) error {
	result, err := h.LocService.GetLastLocation()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, result)
}
