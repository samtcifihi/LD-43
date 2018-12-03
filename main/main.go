package main

import (
	"fmt"
	"github/LD-43/level"
	"github/LD-43/ui"
)

func main() {
	ui.WelcomeUser()

	// Debugging
	myLevel := level.New(3)
	myLevel.Render()

	// Debugging
	gameWon := myLevel.CheckWinCondition()
	fmt.Println("Game Won?:", gameWon)

	fmt.Println("Execution Finished") // Debugging
}
