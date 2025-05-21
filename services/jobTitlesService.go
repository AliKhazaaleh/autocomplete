package services

import (
	"autocomplete/data"
	"autocomplete/pkg/datastructure"
)

var JobTitleTrie *datastructure.Trie
var jobTitles map[int]string

func init() {
	JobTitleTrie = datastructure.NewTrie()
	buildJobTitleTrie()
}

func buildJobTitleTrie() {
	id := 1
	jobTitles = make(map[int]string)
	for _, jobTitle := range data.JobTitleList {
		jobTitles[id] = jobTitle
		JobTitleTrie.Insert(id, jobTitle)
		id++
	}
}

func GetAutoCompleteJobTitles(query string) []string {
	if query == "" {
		return []string{}
	}

	resultIds := JobTitleTrie.Search(query)
	results := make([]string, 0, len(resultIds))
	for _, id := range resultIds {
		results = append(results, jobTitles[id])
	}

	return results
}
