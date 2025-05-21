package services

import (
	"autocomplete/data"
	"autocomplete/pkg/datastructure"
)

var SkillTrie *datastructure.Trie
var skills map[int]string

func init() {
	SkillTrie = datastructure.NewTrie()
	buildSkillTrie()
}

func buildSkillTrie() {
	id := 1
	skills = make(map[int]string)
	for _, skill := range data.SkillList {
		skills[id] = skill
		SkillTrie.Insert(id, skill)
		id++
	}
}

func GetAutoCompleteSkills(query string, limit int) []string {
	if query == "" || limit <= 0 {
		return []string{}
	}

	resultIds := SkillTrie.Search(query)
	results := make([]string, 0, limit)
	for i, id := range resultIds {
		if i >= limit {
			break
		}
		results = append(results, skills[id])
	}

	return results
}
