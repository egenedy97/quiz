package main

import "fmt"

func main() {
	xoBoard := [3][3]string{}
	xPlayer := "X"
	for {
		var vertical int
		var horizontal int
		fmt.Print("Enter a horizontal in range from 1 or 2 or 3\n")
		fmt.Scanln(&horizontal)
		fmt.Print("Enter a vertical in range from 1 or 2 or 3\n")
		fmt.Scanln(&vertical)

		col := vertical - 1
		row := horizontal - 1
		if col < 0 || col > 2 {
			fmt.Print("there is an error in selecting vertical\n")
		}

		if row < 0 || row > 2 {
			fmt.Print("there is an error in selecting horizontal\n")
		}

		if xoBoard[row][col] == "" {
			xoBoard[row][col] = xPlayer
		} else {
			fmt.Println("Invalid entry:", row, col, "value:", xoBoard[row][col])
			continue
		}

		fmt.Println(xoBoard[0])
		fmt.Println(xoBoard[1])
		fmt.Println(xoBoard[2])

		for i := 0; i < 3; i++ {
			// check rows || columns
			if (xoBoard[i][0] == xoBoard[i][1] && xoBoard[i][1] == xoBoard[i][2] && xoBoard[i][2] == xPlayer) ||
				(xoBoard[0][i] == xoBoard[1][i] && xoBoard[1][i] == xoBoard[2][i] && xoBoard[2][i] == xPlayer) {
				fmt.Println("Game ended, winner is player:", xPlayer)
				// we use return to exit both for loops and the app
				return
			}
		}
		if (xoBoard[0][0] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][2] && xoBoard[2][2] == xPlayer) || (xoBoard[0][2] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][0] && xoBoard[2][0] == xPlayer) {
			fmt.Println("Game ended, winner is player:", xPlayer)
			// end the game and exit the app
			// we can use "break" or "return"
			return
		}

		if xPlayer == "X" {
			xPlayer = "O"
		} else {
			xPlayer = "X"
		}
	}

}
