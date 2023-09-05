package chapter01

// IsUnique returns whether a string has all unique characters.
func IsUnique(word string) bool {

	// Using a hashset to store if we've seen a rune
	/*
		seen := map[rune]struct{}{}

		for _, b := range word {
			if _, ok := seen[b]; ok {
				return false
			}
			seen[b] = struct{}{}
		}

		return true
	*/

	// If we asume one byte per rune (ASCII), we can use a 256 long array

	seen := [256]bool{}

	for i := 0; i < len(word); i++ {
		if seen[word[i]] {
			return false
		}
		seen[word[i]] = true
	}

	return true
}

// CheckPermutation (1.2) Given two strings, write a method to decide if one is a permutation of the other
func CheckPermutation(s1, s2 string) bool {

	// Lets get an O(n) solution
	if len(s1) != len(s2) {
		return false
	}

	// put all chars in a map with their frequency
	freqS1 := make(map[byte]int, 0)

	for i := 0; i < len(s1); i++ {
		freqS1[s1[i]]++
	}

	// Iterate over the second string, and reduce the freq for every byte in the map
	// stop early if a char is not found

	for i := 0; i < len(s2); i++ {

		if f1, ok := freqS1[s2[i]]; !ok || f1 == 0 {
			return false
		} else {
			freqS1[s2[i]]--
		}
	}

	return true

	// this solution uses sort so it's O(n log(n))
	/*
		if len(s1) != len(s2) {
			return false
		}

		runes1 := []rune(s1)
		sort.Slice(runes1, func(i, j int) bool {
			return runes1[i] < runes1[j] //sort the string rune
		})

		runes2 := []rune(s2)
		sort.Slice(runes2, func(i, j int) bool {
			return runes2[i] < runes2[j] //sort the string rune
		})

		return string(runes1) == string(runes2)
	*/
}
