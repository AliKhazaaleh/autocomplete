package main

import (
	"autocomplete/controllers"
	"fmt"

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
		return c.File("static/index.html")
	})

	e.GET("/countries", controllers.ActionCountries)
	e.GET("/cities", controllers.ActionCities)

	e.GET("/universities", controllers.ActionUniversities)
	e.GET("/companies", controllers.ActionCompanies)

	e.GET("/job-titles", controllers.ActionJobTitles)
	e.GET("/skills", controllers.ActionSkills)

	e.Logger.Fatal(e.Start(":1323"))
}
