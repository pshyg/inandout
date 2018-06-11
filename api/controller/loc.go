package controller

import (
	"encoding/json"
	"inandout/models"
	"net/http"

	"github.com/labstack/echo"
)

// GetLastLocation return opponent last location
func (h *HTTPLocHandler) GetLastLocation(c echo.Context) error {
	id := c.Param("userid")

	loc, err := h.LocService.GetLastLocation(h.DB, id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, loc.String())
}

// Create create record in DB loc table using request json data
func (h *HTTPLocHandler) Create(c echo.Context) error {
	loc := &models.Location{UserID: c.Param("userid")}
	if err := c.Bind(loc); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.LocService.Create(h.DB, loc)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, result)
}

// GetAllOpponentLoc return all data of opponet
func (h *HTTPLocHandler) GetAllOpponentLoc(c echo.Context) error {
	id := c.Param("userid")

	locs, err := h.LocService.GetAllOpponentLoc(h.DB, id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	returnData, err := json.Marshal(locs)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, string(returnData))
}
