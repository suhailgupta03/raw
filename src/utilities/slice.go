package utilities

func ContainsString(slice []string, val string) bool {
	found := false
	for i := 0; i < len(slice); i++ {
		if slice[i] == val {
			found = true
			break
		}
	}
	return found
}

func IsSubSet(superSlice []string, sliceToCheck []string) bool {
	isSubSet := true

	superSliceMap := make(map[string]bool)
	for i := 0; i < len(superSlice); i++ {
		superSliceMap[superSlice[i]] = true
	}

	for i := 0; i < len(sliceToCheck); i++ {
		_, exists := superSliceMap[sliceToCheck[i]]
		if !exists {
			isSubSet = false
			break
		}
	}

	return isSubSet
}

// AreEqual Checks if the two slices are exactly equal. To be equal the slices
// must have an equal length and should have the same items. The order
// of the items does not matter
func AreEqual(sliceOne []string, sliceTwo []string) bool {
	if len(sliceOne) != len(sliceTwo) {
		return false
	} else {
		sliceOneMap := make(map[string]bool)
		for i := 0; i < len(sliceOne); i++ {
			sliceOneMap[sliceOne[i]] = true
		}

		equal := true
		for i := 0; i < len(sliceOne); i++ { // len(sliceOne) == len(sliceTwo)
			_, exists := sliceOneMap[sliceTwo[i]]
			if !exists {
				equal = false
				break
			}
		}

		return equal
	}
}
