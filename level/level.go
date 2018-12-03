package level

import (
	"bufio"
	"fmt"
	"github/LD-43/mymath"
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

func (l level) CheckWinCondition() bool {
	continueLoop := true
	// Checks all rows for any which are not pairwise coprime
	for i := 0; i < l.m; i++ {
		// Select the next row to check and assign it to b
		b := l.levelSlice[i]

		for continueLoop == true {
			// Pass b into wallthrower, and assign the slice to check if all elements are coprime to a, and teh slice of whatever is leftover to b, to be processed in the next iteration of the loop.
			a, b := wallThrower(b)
			// use SliceCoprime to check a. It will return true if it is pairwise coprime. If it is not, we can return false, otherwise we must continue
			isCoprime, _ := mymath.SliceCoprime(a)
			if isCoprime == false {
				return false
			}

			// Check if b has anything in it. If not, there is nothing more to process and the outer for loop can continue
			if b == nil {
				continueLoop = false
			}
		} // End for
	} // End for

	// Checks all columns for any which are not pairwise coprime
	for i := 0; i < l.n; i++ {
		// Select the next column to check and assign it to b
		b := l.levelSlice[:][i]

		continueLoop = true
		for continueLoop == true {
			// Pass b into wallthrower, and assign the slice to check if all elements are coprime to a, and teh slice of whatever is leftover to b, to be processed in the next iteration of the loop.
			a, b := wallThrower(b)
			// use SliceCoprime to check a. It will return true if it is pairwise coprime. If it is not, we can return false, otherwise we must continue
			isCoprime, _ := mymath.SliceCoprime(a)
			if isCoprime == false {
				return false
			}

			// Check if b has anything in it. If not, there is nothing more to process and the outer for loop can continue
			if b == nil {
				continueLoop = false
			}
		} // End for
	} // End for

	// If both loops have not found any reason to return false, then everything that must be coprime is coprime and we can return true
	return true
}

// Takes in a slice and splits it into a group of positive integers, and everything after that point, and returns both
func wallThrower(inputSlice []int) ([]int, []int) {
	// Beginning, midpoint, and endpoint, roughly speaking
	var a, b, c int

	// Step 0: assign a
	// Step 1: assign b
	// Step 2: assign c
	currentStep := 0

	for i, v := range inputSlice {
		// Check which step we're on
		switch currentStep {
		case 0:
			// If the current element is >= 1, assign its index to a and continue to the next step
			if v >= 1 {
				a = i
				currentStep++
			}
		case 1:
			// If the current element is no longer >= 1, assign its index to b, and the next index to c. Then continue to the last step
			if v < 1 {
				b = i
				c = i + 1
				currentStep++
			}
		case 2:
			if v >= 1 {
				c = len(inputSlice)
			}
		} // End switch
	} // End for

	// If there are no >= 1 numbers in the inputSlice, return something which will fit that situation
	if currentStep == 0 {
		return []int{1}, nil
	}

	// a, b, and c are now assigned values. Return the cleaned up inputSlice and the extra at the end. Note that the second return will be empty if there are no >= 1 elements left in it.
	return inputSlice[a:b], inputSlice[(b + 1):c]
}
