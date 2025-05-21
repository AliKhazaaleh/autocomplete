package main

import (
	"autocomplete/controllers"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	fmt.Println("Server is starting...")
	registerRoutes(e)
	fmt.Println("Server is ready...")

}

func registerRoutes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the Echo server!")
	})

	e.GET("/countries", controllers.ActionCountries)
	e.GET("/cities", controllers.ActionCities)

	e.GET("/universities", controllers.ActionUniversities)
	e.GET("/companies", controllers.ActionCompanies)

	e.GET("/job-titles", controllers.ActionJobTitles)
	e.GET("/skills", controllers.ActionSkills)

	// e.GET("/studies", controllers.ActionStudies)
	// e.GET("/certificates", controllers.ActionCertificates)

	e.Logger.Fatal(e.Start(":1323"))
}
