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
		}
	}

	return isSubSet
}