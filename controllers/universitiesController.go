package controllers

import (
	"autocomplete/pkg/utils"
	"autocomplete/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ActionUniversities(c echo.Context) error {
	query := c.QueryParam("q")
	limit := utils.ParseLimit(c.QueryParam("limit"), 20, 20)
	return c.JSON(http.StatusOK, services.GetAutoCompleteUniversities(query, limit))
}
