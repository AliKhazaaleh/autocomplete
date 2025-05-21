package controllers

import (
	"autocomplete/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ActionUniversities(c echo.Context) error {
	query := c.QueryParam("q")
	return c.JSON(http.StatusOK, services.GetAutoCompleteUniversities(query))
}
