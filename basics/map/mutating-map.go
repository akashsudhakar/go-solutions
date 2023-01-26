package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	out := make(map[string]int)
	for _, word := range words {
		value, exist := out[word]
		if exist {
			out[word] = value + 1
		} else {
			out[word] = 1
		}
	}
	return out
}
