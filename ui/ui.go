package ui

import "fmt"

func WelcomeUser() {
	for i := 0; i <= 5; i++ {
		fmt.Println("")
	}
	fmt.Println("Welcome to this LD 43 game of coprime numbers!")
	fmt.Println("To clear each level you must sacrifice numbers one by one until no two numbers in the same row or column are co-prime with eachother unless separated by a boundary.")
	fmt.Println("If you sacrifice too many numbers, you will fail the level and must begin again.")
}
