package services

import (
	"autocomplete/pkg/datastructure"
	"encoding/json"
	"os"
)

var CityTrie *datastructure.Trie
var cities map[int]string

func init() {
	CityTrie = datastructure.NewTrie()
	buildCityTrie()
}

func buildCityTrie() {
	id := 1
	cities = make(map[int]string)
	file, err := os.Open("data/cities.json")
	if err != nil {
		panic("failed to open cities.json: " + err.Error())
	}
	defer file.Close()

	var countryCities map[string][]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&countryCities); err != nil {
		panic("failed to decode cities.json: " + err.Error())
	}

	for countryName, cityList := range countryCities {
		for _, city := range cityList {
			cities[id] = countryName + " - " + city
			CityTrie.Insert(id, city)
			id++
		}
	}
}

func GetAutoCompleteCities(query string) []string {
	if query == "" {
		return []string{}
	}

	resultIds := CityTrie.Search(query)
	results := make([]string, 0, len(resultIds))
	for _, id := range resultIds {
		results = append(results, cities[id])
	}

	return results
}
