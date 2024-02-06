package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

const (
  rows = 9
  cols = 9
)

func main() {
  // Create a board with empty cells represented by dots.
  board := make([][]rune, rows)
  for i := range board {
    board[i] = make([]rune, cols)
  }

  // Read the Sudoku puzzle from the command line.
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter the Sudoku puzzle (use '.' for empty cells):")
  for i := 0; i < rows; i++ {
    line, _ := reader.ReadString('\n')
    line = strings.TrimSpace(line)
    if len(line) != cols || !isValidInput(line) {
      fmt.Println("Invalid input. Please enter 9 characters ('1'-'9' or '.') per line.")
      return
    }
    for j, c := range line {
      board[i][j] = c
    }
  }

  // Try to solve the board and print the result.
  if solve(board) {
    fmt.Println("Solved Sudoku:")
    print(board)
  } else {
    fmt.Println("No solution found.")
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
  // Find the next empty cell with the fewest possibilities in the board.
  i, j, found := findEmptyWithFewestPossibilities(board)
  // If no empty cell is found, the board is solved.
  if !found {
    return true
  }
  // Try each possible number for the empty cell.
  for _, num := range getPossibleNumbers(board, i, j) {
    // Place the number in the cell.
    board[i][j] = num
    // Recursively try to solve the rest of the board.
    if solve(board) {
      return true
    }
    // Undo the number if it leads to a dead end.
    board[i][j] = '.'
  }
  // Return false if no number works for the cell.
  return false
}

// findEmptyWithFewestPossibilities returns the row, column, and a boolean flag of the next empty cell with the fewest possibilities in the board.
func findEmptyWithFewestPossibilities(board [][]rune) (int, int, bool) {
  minPossibilities := 10
  var minI, minJ int
  found := false
  // Loop through the rows and columns of the board.
  for i := 0; i < rows; i++ {
    for j := 0; j < cols; j++ {
      // If the cell is empty, check the number of possibilities.
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
  // Return the coordinates and true if an empty cell is found, otherwise return zero values and false.
  return minI, minJ, found
}

// getPossibleNumbers returns a slice of possible numbers for a cell.
func getPossibleNumbers(board [][]rune, i, j int) []rune {
  possibilities := make([]rune, 0, 9)
  // Try each number from 1 to 9.
  for num := '1'; num <= '9'; num++ {
    // Check if the number is a valid move for the cell.
    if valid(board, i, j, num) {
      possibilities = append(possibilities, num)
    }
  }
  return possibilities
}

// valid checks if placing a number in a cell is valid according to the Sudoku rules.
func valid(board [][]rune, i, j int, num rune) bool {
  // Check the row and column of the cell.
  for k := 0; k < rows; k++ {
    // If the number already exists in the row or column, return false.
    if board[i][k] == num || board[k][j] == num {
      return false
    }
  }
  // Check the subgrid of the cell.
  subgridSize := 3
  startRow, startCol := subgridSize*(i/subgridSize), subgridSize*(j/subgridSize)
  for i := startRow; i < startRow+subgridSize; i++ {
    for j := startCol; j < startCol+subgridSize; j++ {
      // If the number already exists in the subgrid, return false.
      if board[i][j] == num {
        return false
      }
    }
  }
  // Return true if the number is valid for the cell.
  return true
}

// print prints the board to the standard output.
func print(board [][]rune) {
  for i := 0; i < rows; i++ {
    // Print a newline after every third row.
    if i%3 == 0 && i != 0 {
      fmt.Println()
    }
    for j := 0; j < cols; j++ {
      // Print a space after every third column.
      if j%3 == 0 && j != 0 {
        fmt.Print(" ")
      }
      // Print the cell value.
      fmt.Printf("%c ", board[i][j])
    }
    // Print a newline after every row.
    fmt.Println()
  }
}