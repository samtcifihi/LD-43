package level

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	// Change l.levelSlice to [m][n]
	// Read in data to fill it out.

	// Debugging
	fmt.Println("info: " + l.info)
	fmt.Println("m value: " + strconv.Itoa(l.m))
	fmt.Println("n value: " + strconv.Itoa(l.n))
	fmt.Println("mfmt: " + strconv.Itoa(l.mfmt))
	fmt.Println("nfmt: " + strconv.Itoa(l.nfmt))

	return false
}
