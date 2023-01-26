package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println(Pic(5, 4))
}

func Pic(dx, dy int) [][]uint8 {
	out := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		out[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			out[i][j] = uint8(i * j)
		}
	}
	return out
}
