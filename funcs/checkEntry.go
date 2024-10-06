package b1hangman

import (
	"strings"
)

/** CheckEntry est un checker qui regarde si la proposition correspond à l'une des lettre du mot à chercher
Ce qui renvoi le mot masqué avec les lettres en démasqué.*/
func CheckEntry(input string, toFind string, masked string) string {
	for i, char := range toFind {
		if string(char) == input {
			maskedRunes := []rune(masked)
			maskedRunes[i] = []rune(strings.ToUpper(input))[0]
			masked = string(maskedRunes)
		}
	}
	return masked
}