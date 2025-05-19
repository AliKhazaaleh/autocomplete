package controllers

import (
	"autocomplete/pkg/trie"
	"net/http"

	"github.com/labstack/echo/v4"
)

var companyTrie *trie.Trie

func init() {
	companyTrie = trie.Build(companyList)
}

func ActionCompanies(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusOK, companyList)
	}

	results := companyTrie.Search(query)
	return c.JSON(http.StatusOK, results)
}

var companyList = []string{
	"Apple Inc.",
	"Microsoft Corporation",
	"Amazon.com Inc.",
	"Alphabet Inc. (Google)",
	"Meta Platforms Inc. (Facebook)",
	"Tesla, Inc.",
	"NVIDIA Corporation",
	"Samsung Electronics",
	"Toyota Motor Corporation",
	"Johnson & Johnson",
	"JPMorgan Chase & Co.",
	"Visa Inc.",
	"Walmart Inc.",
	"UnitedHealth Group",
	"Procter & Gamble",
	"Mastercard Incorporated",
	"Bank of America",
	"Netflix, Inc.",
	"Adobe Inc.",
	"Intel Corporation",
}
