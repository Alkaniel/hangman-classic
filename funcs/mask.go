	package b1hangman

	import (
		"math/rand"
	)

	func MaskWord(toFind []rune) []rune {
		var masked []rune
		var notmask []rune
		n := len(toFind)/2 - 1
		for i:= 0; i < n; i++{
			j := rand.Intn(len(toFind)-1)
			notmask = append(notmask, toFind[j])
		}
		for _, letter := range toFind {
			if ContainsForRunes(notmask, letter) {
				masked = append(masked, RuneToUpper(letter))
			} else {
				masked = append(masked, '_')
			}
		}
		return masked
	}

	func ContainsForRunes(letterNotMasked []rune, letter rune) bool {
		for _, r := range letterNotMasked {
			if r == letter {
				return true
			}
		}
		return false
	}

	func RuneToUpper(r rune) rune {
		if r >= 'a' && r <= 'z' {
			r = r - 32
		}
		return r
	}
