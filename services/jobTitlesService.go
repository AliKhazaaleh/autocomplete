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

func GetAutoCompleteJobTitles(query string, limit int) []string {
	if query == "" || limit <= 0 {
		return []string{}
	}

	resultIds := JobTitleTrie.Search(query)
	results := make([]string, 0, limit)
	for i, id := range resultIds {
		if i >= limit {
			break
		}
		results = append(results, jobTitles[id])
	}

	return results
}
