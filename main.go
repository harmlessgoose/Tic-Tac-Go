// Be a 2 player game. I'll be X, and you be O.
// Should print the board between each go using " ", "X", "O"
// Should play it on the command line by inputting coordinates (1,1), (1,2) etc.
// Should stop once someone gets 3 in a row or the board is full!
// Has lots of tests!

package main

import (
	"fmt"
)

type Board struct {
	layout [][]string
}

func (b *Board) printBoard() {
	for _, row := range b.layout {
		fmt.Println(row)
	}
	fmt.Println()
}

func (b *Board) updateBoard(row, col, playerTurn int) {
	if playerTurn == 1 {
		b.layout[row-1][col-1] = "X"
	} else {
		b.layout[row-1][col-1] = "O"
	}
}

func (b *Board) checkWin() bool {
	if checkRows(b.layout) {
		return true
	}
	if checkColumns(b.layout) {
		return true
	}
	if checkDiagonals(b.layout) {
		return true
	}
	return false
}

func main() {
	gameOver := false

	board := Board{
		layout: [][]string{
			{"-", "-", "-"},
			{"-", "-", "-"},
			{"-", "-", "-"},
		},
	}

	playerTurn := 1

	fmt.Println("Welcome to Tic Tac Go!\n")

	for !gameOver {
		// print board
		board.printBoard()

		// get input
		var row, col int

		fmt.Print("Enter two numbers for the row and column (e.g., 1 3): ")
        _, err := fmt.Scan(&row, &col)

		if row < 1 || row > 3 || col < 1 || col > 3 {
			fmt.Println("Invalid input, please enter two integers between 1 and 3.")
			continue
		}

        if err != nil {
            fmt.Println("Invalid input, please enter two integers between 1 and 3.")
            continue
        }

		if board.layout[row-1][col-1] != "-" {
			fmt.Println("That spot is already taken, please choose another.")
			continue
		}

		// update board
		board.updateBoard(row, col, playerTurn)

		// check for win
		if board.checkWin() {
			board.printBoard()
			fmt.Println("You win!")
			break
		}

		// check for draw

		// switch player
		if playerTurn == 1 {
			playerTurn = 2
		} else {
			playerTurn = 1
		}
	}

	
}

func checkRows(board [][]string) bool {
	for i, _ := range board {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != "-" {
			return true
		}
	}
	
	return false
}

// function to check columns
func checkColumns(board [][]string) bool {
	for i, _ := range board {
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != "-" {
			return true
		}
	}
	
	return false
}

// function to check diagonals
func checkDiagonals(board [][]string) bool {
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != "-" {
		return true
	}
	if board[2][0] == board[1][1] && board[1][1] == board[0][2] && board[2][0] != "-" {
		return true
	}
	return false 
}