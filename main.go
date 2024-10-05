package main

import (
	"fmt"
	"os"
)

const size = 9


func solveSudoku(board *[size][size]int) bool {
	row, col := findEmpty(board)
	if row == -1 && col == -1 {
		return true 
	}

	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = 0 
		}
	}

	return false
}


func findEmpty(board *[size][size]int) (int, int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}


func isValid(board *[size][size]int, row, col, num int) bool {
	for i := 0; i < size; i++ {
		if board[row][i] == num || board[i][col] == num ||
			board[row-row%3+i/3][col-col%3+i%3] == num {
			return false
		}
	}
	return true
}


func printBoard(board [size][size]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}


func parseBoard(args []string) ([size][size]int, error) {
	var board [size][size]int

	
	for i := 0; i < size; i++ {
		if len(args[i]) != size {
			return board, fmt.Errorf("Satır %d, %d uzunluğunda. Her satırın tam olarak %d karakter olması gerekiyor.", i+1, len(args[i]), size)
		}
		for j, char := range args[i] {
			if char == '.' {
				board[i][j] = 0
			} else if char >= '1' && char <= '9' {
				board[i][j] = int(char - '0')
			} else {
				return board, fmt.Errorf("Geçersiz karakter '%c' satır %d, sütun %d'de bulundu.", char, i+1, j+1)
			}
		}
	}

	return board, nil
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("error")
		return
	}

	args := os.Args[1:] 
	board, err := parseBoard(args)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	if solveSudoku(&board) {
		printBoard(board)
	} else {
		fmt.Println("Çözüm bulunamadı!")
	}
}
