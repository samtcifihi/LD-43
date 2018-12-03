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

	myLevel.Load()
	myLevel.Render()

	fmt.Println("Execution Finished")
}
