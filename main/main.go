package main

import (
	"bufio"
	"fmt"
	"github/LD-43/level"
	"os"
	"strconv"
	"strings"
)

func main() {
	WelcomeUser()

	newLevel := level.New(3)
	newLevel.Render()

	ShowMenu()

	fmt.Println("Execution Finished") // Debugging
}

func WelcomeUser() {
	for i := 0; i <= 5; i++ {
		fmt.Println("")
	}
	fmt.Println("Welcome to this LD 43 game of coprime numbers!")
	fmt.Println("To clear each level you must sacrifice numbers one by one until no two numbers in the same row or column are co-prime with eachother unless separated by a boundary.")
	fmt.Println("If you sacrifice too many numbers, you will fail the level and must begin again.")
	fmt.Println("Coordinates are entered in the form \"[row] [col]\"")
}

func ShowMenu() {
	// Receive coordinates
	cin := bufio.NewReader(os.Stdin)

	responseInt := -1
	levelWillPlay := false

	for responseInt != 0 {
		levelWillPlay = false
		gettingResponse := true
		for gettingResponse == true {
			// Give some info on the choices
			fmt.Println("type 0 to exit, else type a desired level")

			stringInput, _ := cin.ReadString('\n')
			fmt.Println("stringInput from console:", stringInput) // Debugging

			aInt, _ := strconv.ParseInt(stringInput, 10, 0)
			fmt.Println("assigned to aInt:", aInt) // Debugging

			fmt.Println("assigned to responseInt") // Debugging
			levelWillPlay = true
			responseInt = int(aInt)
			gettingResponse = false
			fmt.Println("responseInt value", responseInt)
		} // End for

		if levelWillPlay == true {
			playLevel(responseInt)
		}
	}
	// Check if they want to exit or choose a level, and if a level, which one.
	// If they choose a non-existant level, say that.
}

func playLevel(chosenLevel int) {
	newLevel := level.New(chosenLevel)
	newLevel.Load()

	gameWon := false

	// Main Game loop
	for gameWon == false {
		newLevel.Render()

		fmt.Println("Please enter two coordinates separated by a space.")

		a, b := 0, 0
		for (a == 0) || (b == 0) {
			sacFailed := true

			for sacFailed == true {
				a, b = getCoordinates()
				sacFailed = newLevel.Sac(a, b)
			}
		}

		gameWon = newLevel.CheckWinCondition()
	}

	fmt.Println("Congratulations, you have won!")
}

func getCoordinates() (int, int) {
	// Receive coordinates
	cin := bufio.NewReader(os.Stdin)
	stringInput, _ := cin.ReadString('\n')
	stringSlice := strings.Split(stringInput, " ")

	if len(stringSlice) != 2 {
		return 0, 0
	}

	outputInt1, err := strconv.Atoi(stringSlice[0])
	if err != nil {
		return 0, 0
	}
	outputInt2, err := strconv.Atoi(stringSlice[1])
	if err != nil {
		return 0, 0
	}

	// Check that they are both ints.
	// If anything is amiss, return 0, 0
	// Else return them
	return outputInt1, outputInt2
}
