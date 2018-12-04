package main

import (
	"fmt"
	"github/LD-43/ui"
	"github/LD-43/level"
)

func main() {
	ui.WelcomeUser()

	newLevel := level.New(3)
	newLevel.Render()

	ui.ShowMenu()

	fmt.Println("Execution Finished") // Debugging
}

