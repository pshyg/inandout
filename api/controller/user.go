package controller

import (
	"inandout/models"
	"net/http"

	"github.com/labstack/echo"
)

// Register register new user
func (h *HTTPUserHandler) Register(c echo.Context) error {
	user := &models.User{ID: c.Param("userid")}
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.UserService.Create(h.DB, user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, result)
}
