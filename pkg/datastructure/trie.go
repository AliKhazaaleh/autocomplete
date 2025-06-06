package datastructure

import (
	"sort"
	"strings"
)

type trieNode struct {
	id       int
	word     string
	children map[rune]*trieNode
	isEnd    bool
}

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{
			children: make(map[rune]*trieNode),
		},
	}
}

func (t *Trie) Insert(id int, word string) {
	node := t.root
	lcWord := strings.ToLower(word)
	for _, ch := range lcWord {
		if node.children[ch] == nil {
			node.children[ch] = &trieNode{
				children: make(map[rune]*trieNode),
			}
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.id = id
	node.word = lcWord
}

func (t *Trie) Search(query string) []int {
	if query == "" {
		return nil
	}

	type match struct {
		id   int
		word string
	}
	var matches []match
	lcQuery := strings.ToLower(query)

	var searchInNode func(node *trieNode)
	searchInNode = func(node *trieNode) {
		if node.isEnd {
			if strings.Contains(node.word, lcQuery) {
				matches = append(matches, match{id: node.id, word: node.word})
			}
		}
		for _, child := range node.children {
			searchInNode(child)
		}
	}

	searchInNode(t.root)
	if len(matches) == 0 {
		return []int{}
	}

	sort.Slice(matches, func(i, j int) bool {
		return strings.Index(matches[i].word, lcQuery) < strings.Index(matches[j].word, lcQuery)
	})

	ids := make([]int, len(matches))
	for i, m := range matches {
		ids[i] = m.id
	}
	return ids
}
