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

	// Nullifies l.levelSlice
	l.levelSlice = nil

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

	return false
}

func (l level) Render() bool {
	// Visual distinction
	for i := 0; i < 6; i++ {
		fmt.Println()
	}

	// Prints info of the level
	fmt.Println(l.info)
	fmt.Println("Sacs:", l.sacs)

	// Prints the column headers
	tempLineString := "  " + "  "
	for i := 1; i <= l.n; i++ {
		tempLineString += l.cellToString(i)
	}
	fmt.Println(tempLineString)

	// Prints a divider
	tempLineString = "  " + "  "
	for i := 1; i <= l.n; i++ {
		for j := l.nfmt; j >= 0; j-- {
			tempLineString += "-"
		}
	}
	fmt.Println(tempLineString)

	// Prints the remaining lines
	for i := 0; i < l.m; i++ {
		// Print Row header
		tempnfmt := l.nfmt
		l.nfmt = 2
		tempLineString = l.cellToString(i + 1)
		l.nfmt = tempnfmt

		// Print a divider
		tempLineString += "|"

		// Print the values of the cells in the row
		for j := 0; j < l.n; j++ {
			tempLineString += l.cellToString(l.levelSlice[i][j])
		}

		// Actually outputs tempLineString
		fmt.Println(tempLineString)

		// Prints any extra blank lines (number extra == mfmt) below
		for mf := l.mfmt; mf > 0; mf-- {
			tempLineString = "  " + " |"
		}
		fmt.Println(tempLineString)
	}

	return false
}

func (l *level) Sac(i int, j int) bool {
	// Check if i and j are within bounds. If not, return true and change nothing
	if ((i - 1) < l.m) && ((j - 1) < l.n) {
		// Do nothing
	} else {
		return true
	}

	// Sac the appropriate cell.
	l.levelSlice[i-1][j-1] = 0

	// Increment Sac Count
	l.sacs++
	return false
}

func (l level) cellToString(input int) string {
	outputString := ""

	if input == 0 {
		// Adds nfmt + 1 spaces to outputString
		outputString = " "
		for i := 0; i < l.nfmt; i++ {
			outputString += " "
		}
	} else if input > 0 {
		// Normal input
		inputString := strconv.Itoa(input)

		// inputLen := strings.LastIndexAny(inputString, "0123456789") + 1
		inputLen := len(inputString)

		if inputLen <= l.nfmt {
			// Adds an extra to put space between entries
			spacesToAdd := l.nfmt - inputLen + 1
			for i := spacesToAdd; i > 0; i-- {
				outputString += " "
			} // End for
			outputString += inputString
		} else {
			// Simple case
			outputString = inputString
		}
	} else {
		switch input {
		case -1: // Wall
			// Adds nfmt asterisks and a leading space to outputString
			outputString = " "
			for i := 0; i < l.nfmt; i++ {
				outputString += "*"
			}
		} // End switch
	} // End if-else

	return outputString
}
