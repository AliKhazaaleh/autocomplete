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

func GetAutoCompleteCompanies(query string, limit int) []string {
	if query == "" || limit <= 0 {
		return []string{}
	}

	resultIds := CompanyTrie.Search(query)
	results := make([]string, 0, limit)
	for i, id := range resultIds {
		if i >= limit {
			break
		}
		results = append(results, companies[id])
	}

	return results
}
