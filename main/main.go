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

	fmt.Println("--------------")

	myLevel.Load()
	myLevel.Render()
}
