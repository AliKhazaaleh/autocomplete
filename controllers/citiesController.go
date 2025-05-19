package controllers

import (
	"autocomplete/pkg/trie"
	"net/http"

	"github.com/labstack/echo/v4"
)

var cityTrie *trie.Trie

func init() {
	cityTrie = trie.Build(cityList)
}

func ActionCities(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusOK, cityList)
	}

	results := cityTrie.Search(query)
	return c.JSON(http.StatusOK, results)
}

var cityList = []string{
	"New York",
	"Los Angeles",
	"Chicago",
	"Houston",
	"Phoenix",
	"Philadelphia",
	"San Antonio",
	"San Diego",
	"Dallas",
	"San Jose",
	"Tokyo",
	"Delhi",
	"Shanghai",
	"São Paulo",
	"Mumbai",
	"Beijing",
	"Cairo",
	"Dhaka",
	"Mexico City",
	"London",
	"Paris",
	"Berlin",
	"Madrid",
	"Rome",
	"Toronto",
	"Sydney",
	"Singapore",
	"Dubai",
}
