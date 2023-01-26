package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

func main() {
	input := `The quick brown fox jumped over the lazy dog's bowl.
		The dog was angry with the fox for considering him lazy.`

	words := clean(input)
	sortWords(words)

	// 	the 4
	//  fox 2
	//  lazy 2
	//  angry 1
	//  bowl 1
	//  brown 1
	//  considering 1
	//  dog 1
	//  dogs 1
	//  for 1
}

type wordCount struct {
	word  string
	count int
}

type wordSort []wordCount

func (a wordSort) Len() int      { return len(a) }
func (a wordSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a wordSort) Less(i, j int) bool {
	if a[i].count != a[j].count {
		return a[i].count > a[j].count
	}
	return a[i].word < a[j].word
}

func sortWords(input []string) {
	wordMap := make(map[string]int)
	for _, word := range input {
		wordMap[word] = wordMap[word] + 1
	}

	var wordCountSlice []wordCount
	for key, value := range wordMap {
		wordCountSlice = append(wordCountSlice, wordCount{key, value})
	}

	sort.Sort(wordSort(wordCountSlice))

	for _, elem := range wordCountSlice[:10] {
		fmt.Println(elem.word, elem.count)
	}
}

func clean(str string) []string {
	words := strings.Fields(strings.ToLower(str))

	// Remove all non letter characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	for ii, elem := range words {
		words[ii] = reg.ReplaceAllString(elem, "")
	}

	return words
}
