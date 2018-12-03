package mymath

import (
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
		myBigIntc = myBigIntc.GCD(nil, nil, myBigInta, myBigIntb)
		resultCompare := myBigIntc.Cmp(myBigInt1)

		// Check if they are coprime and if so, return true, else, return false.
		if resultCompare == 0 {
			return true, false
		} else {
			return false, false
		}
	}

	return false, true
}

// Tests if any of the numbers in an array are coprime to eachother and returns the answer as well as an error flag.
func SliceCoprime(inputSlice []int) (bool, bool) {
	// Test if the slice is big enough to check for coprimacy
	if len(inputSlice) == 0 {
		return false, true
	} else if len(inputSlice) == 1 {
		return true, false
	} else {
		return reduceSlice(inputSlice), false
	}
}

// Recursive function which accepts an inputSlice and returns whether or not it is true that all elements are coprime to eachother
func reduceSlice(inputSlice []int) bool {
	// Check if base case is reached. If so, Everything is coprime to everything else.
	if len(inputSlice) <= 1 {
		return true
	} else {
		for i, v := range inputSlice {
			// Skip the first iteration
			if i == 0 {
				// Do nothing
			} else {
				// Check if v is coprime to i == 0. If so, do nothing, else return false
				isCoprime, _ := coprime(v, inputSlice[0])
				if isCoprime {
					// Do nothing
				} else {
					// End the recursion as the original slice is not pairwise coprime
					return false
				} // End if-else
			} // End if-else
		} // End for
		// Remove first value from inputSlice and recursively call reduceSlice
		outputSlice := inputSlice[1:]
		return reduceSlice(outputSlice)
	} // End if-else
}
