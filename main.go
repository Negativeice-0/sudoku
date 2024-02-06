package main

import (
	"fmt"
	"os"
)

const (
	rows = 9
	cols = 9
)

func main() {
	// Check if the correct number of arguments is provided.
	if len(os.Args) != rows+1 {
		fmt.Println("Error")
		return
	}

	// Create a board with cells filled according to the command line arguments.
	board := make([][]rune, rows)
	for i := range board {
		board[i] = make([]rune, cols)
		line := os.Args[i+1]
		// Check if the input string is valid.
		if len(line) != cols || !isValidInput(line) {
			fmt.Println("Error")
			return
		}
		for j, c := range line {
			board[i][j] = c
		}
	}

	// Try to solve the board and print the result.
	if solve(board) {
		print(board)
	} else {
		fmt.Println("Error")
	}
}

// isValidInput checks if the input string is valid (i.e., contains only '1'-'9' or '.').
func isValidInput(input string) bool {
	for _, c := range input {
		if c < '1' || c > '9' && c != '.' {
			return false
		}
	}
	return true
}

// solve tries to fill the board with valid numbers and returns true if successful.
func solve(board [][]rune) bool {
	i, j, found := findEmptyWithFewestPossibilities(board)
	if !found {
		return true
	}
	for _, num := range getPossibleNumbers(board, i, j) {
		board[i][j] = num
		if solve(board) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

// findEmptyWithFewestPossibilities returns the row, column, and a boolean flag of the next empty cell with the fewest possibilities in the board.
func findEmptyWithFewestPossibilities(board [][]rune) (int, int, bool) {
	minPossibilities := 10
	var minI, minJ int
	found := false
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '.' {
				possibilities := getPossibleNumbers(board, i, j)
				if len(possibilities) < minPossibilities {
					minPossibilities = len(possibilities)
					minI = i
					minJ = j
					found = true
				}
			}
		}
	}
	return minI, minJ, found
}

// getPossibleNumbers returns a slice of possible numbers for a cell.
func getPossibleNumbers(board [][]rune, i, j int) []rune {
	possibilities := make([]rune, 0, 9)
	for num := '1'; num <= '9'; num++ {
		if valid(board, i, j, num) {
			possibilities = append(possibilities, num)
		}
	}
	return possibilities
}

// valid checks if placing a number in a cell is valid according to the Sudoku rules.
func valid(board [][]rune, i, j int, num rune) bool {
	for k := 0; k < rows; k++ {
		if board[i][k] == num || board[k][j] == num {
			return false
		}
	}
	subgridSize := 3
	startRow, startCol := subgridSize*(i/subgridSize), subgridSize*(j/subgridSize)
	for i := startRow; i < startRow+subgridSize; i++ {
		for j := startCol; j < startCol+subgridSize; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
}

// print prints the board to the standard output.
func print(board [][]rune) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Println()
	}
}

//  go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...."
// "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
// | cat -e
//3 9 6 2 4 5 7 8 1$
//Error$
// gol.21.6 linux/amd64
