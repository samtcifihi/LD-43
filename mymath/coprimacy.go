package mymath

import (
	"fmt"
	"math/big"
)

// Tests if two numbers are coprime to eachother and returns the answer as well as an error flag
func coprime(a, b int) (bool, bool) {
	myBigInta := big.NewInt(int64(a))
	myBigIntb := big.NewInt(int64(b))
	myBigIntc := big.NewInt(int64(1))
	myBigInt1 := big.NewInt(int64(1))

	// Checks for valid input
	if (a >= 1) && (b >= 1) {
		fmt.Println("Passed Validation") // Debugging
		myBigIntc = myBigIntc.GCD(nil, nil, myBigInta, myBigIntb)
		fmt.Println(myBigIntc) // Debugging
		fmt.Println(myBigInt1) // Debugging
		resultCompare := myBigIntc.Cmp(myBigInt1)

		// Check if they are coprime and if so, return true, else, return false.
		if resultCompare == 0 {
			fmt.Println("passed coprime == true test") // Debugging
			return true, false
		} else {
			fmt.Println("failed coprime == true test") // Debugging
			return false, false
		}
	}

	fmt.Println("Missed coprime == true test altogether") // Debugging
	return false, true
}

// Tests if any of the numbers in an array are coprime to eachother and returns the answer as well as an error flag.
func ArrCoprime(inputSlice []int) (bool, bool) {
	// Use "Range" to know when to stop the nested for loops
	return false, false
}
