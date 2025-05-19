package controllers

import (
	"autocomplete/pkg/trie"
	"net/http"

	"github.com/labstack/echo/v4"
)

var skillTrie *trie.Trie

func init() {
	skillTrie = trie.Build(skillList)
}

func ActionSkills(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusOK, skillList)
	}

	results := skillTrie.Search(query)
	return c.JSON(http.StatusOK, results)
}

var skillList = []string{
	"JavaScript",
	"Python",
	"Java",
	"C++",
	"Go",
	"Ruby",
	"TypeScript",
	"SQL",
	"React",
	"Angular",
	"Vue.js",
	"Node.js",
	"Docker",
	"Kubernetes",
	"AWS",
	"Azure",
	"Git",
	"Machine Learning",
	"Deep Learning",
	"Data Analysis",
	"Project Management",
	"Agile",
	"Scrum",
	"Communication",
	"Leadership",
	"Problem Solving",
	"Team Collaboration",
	"Critical Thinking",
}
