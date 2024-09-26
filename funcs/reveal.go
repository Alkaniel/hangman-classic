package b1hangman

import (
	"math/rand"
	"strings"
)

func MaskWord(toFind string) string {
	masked := ""
	var notmask []rune
	n := len(toFind)/2 - 1
	for i:= 0; i == n; i++{
		j := rand.Intn(len(toFind)-1)
		notmask = append(notmask, []rune(toFind)[j])
	}
	for _, letter := range toFind {
		if ContainsForRunes(notmask, letter) {
			masked += string(letter) + " "
		} else {
			masked += "_ "
		}
	}
	return strings.TrimSpace(masked)
}

func ContainsForRunes(letterNotMasked []rune, letter rune) bool {
	for _, r := range letterNotMasked {
		if r == letter {
			return true
		}
	}
	return false
}
