package main

import (
	"fmt"
	"github/LD-43/level"
	"github/LD-43/mymath"
	"github/LD-43/ui"
)

func main() {
	ui.WelcomeUser()

	// Debugging
	myLevel := level.New(1)
	myLevel.Render()

	// Debugging
	mySlice := []int{2, 6, 3}
	myBool1, myBool2 := mymath.SliceCoprime(mySlice)
	fmt.Println("SliceCoprime result:", myBool1, myBool2)

	fmt.Println("Execution Finished") // Debugging
}
