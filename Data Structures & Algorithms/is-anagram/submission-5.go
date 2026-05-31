import "slices"

func isAnagram(s string, t string) bool {
	
	if len(s) != len(t) {
		return false
	}

	runesT := []rune(t)
	runesS := []rune(s)

	slices.Sort(runesT)
	slices.Sort(runesS)

	sortedS := string(runesS)
	sortedT := string(runesT)

	if sortedS == sortedT {
		return true
	} 
	return false
}

