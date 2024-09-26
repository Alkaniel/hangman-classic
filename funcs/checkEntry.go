package b1hangman

func CheckEntry(letter byte, toFind string, masked string) string{
	maskedAsRune := []rune(masked)
	for i, char := range toFind {
		if char == rune(letter) {
			maskedAsRune[i] = rune(letter)
		}
	}
	return string(maskedAsRune)
}