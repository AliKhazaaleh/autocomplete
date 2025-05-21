package services

import (
	"autocomplete/data"
	"autocomplete/pkg/datastructure"
)

var UniversityTrie *datastructure.Trie
var universities map[int]string

func init() {
	UniversityTrie = datastructure.NewTrie()
	buildUniversityTrie()
}

func buildUniversityTrie() {
	id := 1
	universities = make(map[int]string)
	for _, university := range data.UniversityList {
		universities[id] = university
		UniversityTrie.Insert(id, university)
		id++
	}
}

func GetAutoCompleteUniversities(query string, limit int) []string {
	if query == "" || limit <= 0 {
		return []string{}
	}

	resultIds := UniversityTrie.Search(query)
	results := make([]string, 0, limit)
	for i, id := range resultIds {
		if i >= limit {
			break
		}
		results = append(results, universities[id])
	}

	return results
}
