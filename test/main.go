package main

import (
	"fmt"
	"os"
)

func main() {
	// check if the arguments are 9
	if len(os.Args) != 10 {
		fmt.Println("Error: number of arguments must be 9")
		return
	}

	// create a 2d array to store the sudoku arguments
	sudoku := make([][]int, 9)
	for i := range sudoku {
		sudoku[i] = make([]int, 9)
		for j, char := range os.Args[i+1] {
			// check if its a digit or a dot, if a dot change it to 0
			if char == '.' {
				sudoku[i][j] = 0
			} else {
				sudoku[i][j] = int(char - '0')
			}
		}
		// print unsolved sudoku to check
		// fmt.Println(sudoku[i])
	}

	// solve the sudoku
	if solveSudoku(sudoku) {
		// if solved, print the solved sudoku
		for _, row := range sudoku {
			for _, num := range row {
				fmt.Print(num, " ")
			}
			fmt.Println()
		}
	} else {
		// cant be solved
		fmt.Println("Error: there is no solution")
	}
}

// solving the Sudoku
func solveSudoku(sudoku [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// check if the sudoku at this position is empty (0)
			if sudoku[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					// looks for a valid number to place it
					if isValid(sudoku, row, col, num) {
						sudoku[row][col] = num
						// if the number is valid, check the next position
						if solveSudoku(sudoku) {
							// if all positions are filled, return true
							return true
						}
						// if the number is not valid, reset the position to 0
						sudoku[row][col] = 0
					}
				}
				// if no number is valid
				return false
			}
		}
	}
	// if all positions are filled return true
	return true
}

// check if placing num at board[row][col] is valid
func isValid(sudoku [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		// check if the number is already in the row, column or the 3x3 box
		if sudoku[row][i] == num || sudoku[i][col] == num || !check3x3(sudoku, row-row%3, col-col%3, num) {
			// if so return false to skip this number
			return false
		}
	}
	// else return true
	return true
}

// check if placing num at board[row][col] is valid in the 3x3 box
func check3x3(sudoku [][]int, row, col, num int) bool {
	// for each row and column in the 3x3 box
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// check if the number is already in the 3x3 box
			if sudoku[row+i][col+j] == num {
				return false
			}
		}
	}
	return true
}
