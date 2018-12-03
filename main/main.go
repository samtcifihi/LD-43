package main

import (
	"fmt"
	"github/LD-43/level"
	"github/LD-43/ui"
)

func main() {
	ui.WelcomeUser()

	myLevel := level.New(1)
	myLevel.Render()
	myLevel.Sac(1, 2)
	myLevel.Render()

	myLevel = level.New(2)
	myLevel.Render()

	myLevel = level.New(3)
	myLevel.Render()

	fmt.Println("Execution Finished")
}
