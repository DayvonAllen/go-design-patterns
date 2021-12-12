package template_method

import "fmt"

func PlayGameFunctional(start, takeTurn func(),
	haveWinner func() bool,
	winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.\n", winningPlayer())
}

func StartFunctionalChess() {
	turn, maxTurns, currentPlayer := 1, 10, 0

	start := func() {
		fmt.Println("Starting a game of chess.")
	}

	takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d\n",
			turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGameFunctional(start, takeTurn, haveWinner, winningPlayer)
}
