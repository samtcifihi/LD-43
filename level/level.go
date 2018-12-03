package level

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type level struct {
	// The file path to the level.
	levelPath string

	// Number of sacrifices made since last level reset.
	sacs int

	// String which will print above each level explaining the goal.
	info string

	// Specifies the number of rows (m) and number of columns (n) in the level.
	m int
	n int

	// Specifies how many blank lines will be added above each line of the level (mfmt)  and how many characters each number of the level will occupy (nfmt).
	mfmt int
	nfmt int

	// The 2-D slice to hold the level
	// var levelSlice := make([][]int, 1)
	levelSlice [][]int
}

// Constructor
func New(levelID int) *level {
	newLevel := new(level)
	newLevel.levelPath = "lvl" + strconv.Itoa(levelID) + ".txt"

	newLevel.Load()

	fmt.Println("nfmt in new after Load has run:", newLevel.nfmt) // Debugging

	return newLevel
}

func (l *level) Load() bool {
	// Reset sacs counter to 0
	l.sacs = 0

	// Initialize streamreader
	file, err := os.Open(l.levelPath)
	if err != nil {
		return true
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Read in info from first line of levelPath.
	scanner.Scan()
	l.info = scanner.Text()

	// Read in m from second line of levelPath.
	scanner.Scan()
	l.m, err = strconv.Atoi(scanner.Text())

	// Read in n from third line of levelPath.
	scanner.Scan()
	l.n, err = strconv.Atoi(scanner.Text())

	// Read in mfmt from fourth line of levelPath.
	scanner.Scan()
	l.mfmt, err = strconv.Atoi(scanner.Text())

	// Read in nfmt from fifth line of levelPath.
	scanner.Scan()
	l.nfmt, err = strconv.Atoi(scanner.Text())
	fmt.Println("l.nfmt: ", l.nfmt) // Debugging

	// Nullifies l.levelSlice
	l.levelSlice = nil
	fmt.Println("l.nfmt After nilling levelSlice: ", l.nfmt) // Debugging

	// Read in data to fill it out.
	for i := 0; i < l.m; i++ {
		scanner.Scan() // Reads in the next line

		// Splits the output of the scanner into a slice, space delineated
		tempStringRowSlice := strings.Split(scanner.Text(), " ")
		tempIntRowSlice := []int{}

		for j := 0; j < l.n; j++ {

			tempIntVal := 0
			tempIntVal, err = strconv.Atoi(tempStringRowSlice[j])
			if err != nil {
				return true
			}

			tempIntRowSlice = append(tempIntRowSlice, tempIntVal)
		} // End for

		l.levelSlice = append(l.levelSlice, tempIntRowSlice)
	} // End for
	fmt.Println("l.nfmt at the end of Load: ", l.nfmt) // Debugging

	return false
}

func (l level) Render() bool {
	fmt.Println("Render reached") // Debugging
	fmt.Println("l.nfmt at the start of Render: ", l.nfmt)
	testString := l.cellToString(234)
	fmt.Println("testString: \"" + testString + "\"") // Debugging
	testString = l.cellToString(48)
	fmt.Println("testString: \"" + testString + "\"") // Debugging
	testString = l.cellToString(4)
	fmt.Println("testString: \"" + testString + "\"") // Debugging
	return false
}

func (l level) cellToString(input int) string {
	outputString := ""

	if input == 0 {
		// Adds nfmt spaces to outputString
		for i := 0; i < l.nfmt; i++ {
			outputString += " "
		}
	} else if input > 0 {
		// Normal input
		inputString := strconv.Itoa(input)

		fmt.Println("input >0 reached.") // Debugging
		inputLen := strings.LastIndexAny(inputString, "0123456789") + 1
		fmt.Println("inputLen: ", strconv.Itoa(inputLen)) // Debugging
		fmt.Println("l.nfmt: ", strconv.Itoa(l.nfmt))     // Debugging

		if inputLen < l.nfmt {
			spacesToAdd := l.nfmt - inputLen
			for i := spacesToAdd; i > 0; i-- {
				outputString += " "
			} // End for
			fmt.Println("outputString: \"" + outputString + "\"")
			outputString += inputString
			fmt.Println("inputString: \"" + inputString + "\"")
		} else {
			// Simple case
			outputString = inputString
		}
	} else {
		switch input {
		case -1: // Wall
			// Adds nfmt asterisks to outputString
			for i := 0; i < l.nfmt; i++ {
				outputString += "*"
			}
		} // End switch
	} // End if-else

	fmt.Println("Exiting cellToString") // Debugging
	return outputString
}
