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

func GetAutoCompleteSkills(query string) []string {
	if query == "" {
		return []string{}
	}

	resultIds := SkillTrie.Search(query)
	results := make([]string, 0, len(resultIds))
	for _, id := range resultIds {
		results = append(results, skills[id])
	}

	return results
}
