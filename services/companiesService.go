package services

import (
	"autocomplete/data"
	"autocomplete/pkg/datastructure"
)

var CompanyTrie *datastructure.Trie
var companies map[int]string

func init() {
	CompanyTrie = datastructure.NewTrie()
	buildCompanyTrie()
}

func buildCompanyTrie() {
	id := 1
	companies = make(map[int]string)
	for _, company := range data.CompanyList {
		companies[id] = company
		CompanyTrie.Insert(id, company)
		id++
	}
}

func GetAutoCompleteCompanies(query string) []string {
	if query == "" {
		return []string{}
	}

	resultIds := CompanyTrie.Search(query)
	results := make([]string, 0, len(resultIds))
	for _, id := range resultIds {
		results = append(results, companies[id])
	}

	return results
}
