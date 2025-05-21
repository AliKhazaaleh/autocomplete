package controllers

import (
	"autocomplete/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ActionJobTitles(c echo.Context) error {
	query := c.QueryParam("q")
	return c.JSON(http.StatusOK, services.GetAutoCompleteJobTitles(query))
}
