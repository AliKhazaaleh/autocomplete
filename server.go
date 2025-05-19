package main

import (
	"autocomplete/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the Echo server!")
	})

	// Countries routes
	e.GET("/countries", controllers.GetCountrySuggestions)

	e.Logger.Fatal(e.Start(":1323"))
}
