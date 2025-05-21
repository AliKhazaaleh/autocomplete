package services

import (
	"autocomplete/data"
	"autocomplete/pkg/datastructure"
)

var CountryTrie *datastructure.Trie
var countries map[int]string

func init() {
	CountryTrie = datastructure.NewTrie()
	buildCountryTrie()
}

func buildCountryTrie() {
	id := 1
	countries = make(map[int]string)
	for _, country := range data.CountryList {
		countries[id] = country
		CountryTrie.Insert(id, country)
		id++
	}
}

func GetAutoCompleteCountries(query string, limit int) []string {
	if query == "" || limit <= 0 {
		return []string{}
	}

	resultIds := CountryTrie.Search(query)
	results := make([]string, 0, limit)
	for i, id := range resultIds {
		if i >= limit {
			break
		}
		results = append(results, countries[id])
	}

	return results
}
