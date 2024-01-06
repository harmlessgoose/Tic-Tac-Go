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


	for !gameOver {
		// print board
		printBoard(board)

		// get input

		// update board

		// check for win
		gameOver = checkWin()

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
}