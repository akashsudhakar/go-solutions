package main

import (
	"errors"
)

func reverseString(input string) (string, error) {
	if input == "" {
		return input, errors.New("empty input")
	}
	in := []rune(input)
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return string(in), nil
}
