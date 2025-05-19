package controllers

import (
	"autocomplete/pkg/trie"
	"net/http"

	"github.com/labstack/echo/v4"
)

var jobTitleTrie *trie.Trie

func init() {
	jobTitleTrie = trie.Build(jobTitleList)
}

func ActionJobTitles(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusOK, jobTitleList)
	}

	results := jobTitleTrie.Search(query)
	return c.JSON(http.StatusOK, results)
}

var jobTitleList = []string{
	"Software Engineer",
	"Project Manager",
	"Data Scientist",
	"Product Manager",
	"Business Analyst",
	"Marketing Manager",
	"Sales Representative",
	"Human Resources Manager",
	"Financial Analyst",
	"Operations Manager",
	"Systems Administrator",
	"UX Designer",
	"DevOps Engineer",
	"Web Developer",
	"Mobile Developer",
	"Cloud Architect",
	"Network Engineer",
	"Security Engineer",
	"Machine Learning Engineer",
	"Full Stack Developer",
}
