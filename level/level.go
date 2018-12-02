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

func New(levelID int) *level {
	newLevel := new(level)
	newLevel.levelPath = "lvl" + strconv.Itoa(levelID) + ".txt"

	newLevel.Load()

	return newLevel
}

func (l level) Load() bool {
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

	// Debugging
	fmt.Println("Just before scanning in the level proper.")

	// Nullifies l.levelSlice
	l.levelSlice = nil
	// Reslices l.levelSlice to [m][n]
	// l.levelSlice = l.levelSlice[:0][:0]

	// Debugging
	fmt.Println("Just after Slicing to size 0")

	// l.levelSlice = l.levelSlice[:l.m][:l.n]

	// Debugging
	fmt.Println("Just after reslicing l.levelSlice.")

	// Read in data to fill it out.
	for i := 0; i < l.m; i++ {
		scanner.Scan() // Reads in the next line

		// Splits the output of the scanner into a slice, space delineated
		tempStringRowSlice := strings.Split(scanner.Text(), " ")
		tempIntRowSlice := []int{}

		fmt.Println("In i loop...")                             // Debugging
		fmt.Println("scanner.Text(): ", scanner.Text())         // Debugging
		fmt.Println("tempStringRowSlice: ", tempStringRowSlice) // Debugging
		fmt.Println("-----")                                    // Debugging

		for j := 0; j < l.n; j++ {
			fmt.Println("Entered j loop...") // Debugging

			tempIntVal := 0
			tempIntVal, err = strconv.Atoi(tempStringRowSlice[j])
			if err != nil {
				return true
			}

			tempIntRowSlice = append(tempIntRowSlice, tempIntVal)

			fmt.Println("tempIntVal: ", strconv.Itoa(tempIntVal))
			fmt.Println("tempIntRowSlice: ", tempIntRowSlice) // Debugging
			fmt.Println("Successfully appended in j loop")
			fmt.Println("-----") // Debugging
		} // End for

		fmt.Println("In i loop after j loop...")                  // Debugging
		fmt.Println("l.levelSlice before append: ", l.levelSlice) // Debugging
		l.levelSlice = append(l.levelSlice, tempIntRowSlice)
		fmt.Println("l.levelSlice after append: ", l.levelSlice) // Debugging
	} // End for

	// Debugging
	fmt.Println("info: ", l.info)
	fmt.Println("m value: ", strconv.Itoa(l.m))
	fmt.Println("n value: ", strconv.Itoa(l.n))
	fmt.Println("mfmt: ", strconv.Itoa(l.mfmt))
	fmt.Println("nfmt: ", strconv.Itoa(l.nfmt))

	return false
}
