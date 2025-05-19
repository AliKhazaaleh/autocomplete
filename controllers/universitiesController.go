package controllers

import (
	"autocomplete/pkg/trie"
	"net/http"

	"github.com/labstack/echo/v4"
)

var universityTrie *trie.Trie

func init() {
	universityTrie = trie.Build(universityList)
}

func ActionUniversities(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusOK, universityList)
	}

	results := universityTrie.Search(query)
	return c.JSON(http.StatusOK, results)
}

var universityList = []string{
	"Harvard University",
	"Massachusetts Institute of Technology",
	"Stanford University",
	"University of Cambridge",
	"University of Oxford",
	"California Institute of Technology",
	"Princeton University",
	"Columbia University",
	"Yale University",
	"University of Chicago",
	"University of California, Berkeley",
	"Imperial College London",
	"ETH Zurich",
	"University of Toronto",
	"National University of Singapore",
	"Peking University",
	"Tsinghua University",
	"University of Melbourne",
	"University of Munich",
	"University of Tokyo",
}
