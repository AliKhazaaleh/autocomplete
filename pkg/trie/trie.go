package trie

import (
	"sort"
	"strings"
)

type trieNode struct {
	children map[rune]*trieNode
	isEnd    bool
	word     string
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

func (t *Trie) insert(word string) {
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
	node.word = word
}

func Build(words []string) *Trie {
	trie := NewTrie()
	for _, word := range words {
		trie.insert(word)
	}
	return trie
}

// Search finds all words that contain the query string and sorts them by relevance
func (t *Trie) Search(query string) []string {
	if query == "" {
		return nil
	}

	var matches []string
	lcQuery := strings.ToLower(query)

	var searchInNode func(node *trieNode, currentWord string)
	searchInNode = func(node *trieNode, currentWord string) {
		if node.isEnd {
			lcWord := strings.ToLower(node.word)
			if strings.Contains(lcWord, lcQuery) {
				matches = append(matches, node.word)
			}
		}
		for ch, child := range node.children {
			searchInNode(child, currentWord+string(ch))
		}
	}

	searchInNode(t.root, "")
	// Sort matches by the index of the query in each word (lower index first)
	if len(matches) == 0 {
		return []string{}
	}

	sort.Slice(matches, func(i, j int) bool {
		li := strings.ToLower(matches[i])
		lj := strings.ToLower(matches[j])
		return strings.Index(li, lcQuery) < strings.Index(lj, lcQuery)
	})

	return matches

}
