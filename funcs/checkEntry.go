package b1hangman

func CheckEntry(input rune, toFind []rune, masked []rune) []rune {
	for i, char := range toFind {
		if char == input {
			masked[i] = RuneToUpper(input)
		}
	}
	return masked
}