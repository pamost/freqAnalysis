package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func countWords(s string) {

	type storage struct {
		Key   string
		Value int
	}

	// Exclude special characters
	var repl = regexp.MustCompile(`[[:punct:]]`)
	str := repl.ReplaceAllString(s, "")

	// Convert to lowercase and remove extra spaces
	words := strings.Fields(strings.ToLower(str))

	// Count the number of words and fill in the map
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	// Map sorting
	var mapWords []storage
	for k, v := range counts {
		mapWords = append(mapWords, storage{k, v})
	}
	sort.Slice(mapWords, func(i, j int) bool {
		return mapWords[i].Value > mapWords[j].Value
	})

	// Output the first 10 values
	for i, w := range mapWords {
		i++
		if i > 10 {
			break
		}
		fmt.Println(w.Key, w.Value)
	}
}

func main() {

	// Read file to byte slice
	s, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	countWords(string(s))
}