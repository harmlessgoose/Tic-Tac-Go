// Be a 2 player game. I'll be X, and you be O.
// Should print the board between each go using " ", "X", "O"
// Should play it on the command line by inputting coordinates (1,1), (1,2) etc.
// Should stop once someone gets 3 in a row or the board is full!
// Has lots of tests!

package main

import (
	"fmt"
)

func main() {
	gameOver := false

	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	playerturn := 1

	fmt.Println("Welcome to Tic Tac Go!\n")

	for !gameOver {
		// print board
		printBoard(board)

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

		if board[row-1][col-1] != "-" {
			fmt.Println("That spot is already taken, please choose another.")
			continue
		}

		// update board
		if playerturn == 1 {
			board[row-1][col-1] = "X"
		} else {
			board[row-1][col-1] = "O"
		}

		// check for win

		// check for draw

		// switch player
	}

	
	fmt.Println("You win!")
}

func checkWin() bool {
	return true
}

func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println()
}