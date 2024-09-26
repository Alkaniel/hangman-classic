package b1hangman

import (
	"math/rand"
	"string"
	"time"
)

func MaskWord(toFind string) string {
	rand.Seed(time.Now().UnixNano())

	n := len(toFind)/2 - 1
	if n < 1 {
		n = 1
	}

	var notmask []rune

	for i := 0; i < n; i++ {
		j := rand.Intn(len(toFind))
		notmask = append(notmask, []rune(toFind)[j])
	}

	masked := make([]rune, len(toFind))

	for i, letter := range toFind {
		if containsRune(notmask, letter) {
			masked[i] = letter
		} else {
			masked[i] = '_'
		}
	}

	maskedWord := string(masked)

	return maskedWord
}

func containsRune(slice []rune, r rune) bool {
	for _, s := range slice {
		if s == r {
			return true
		}
	}
	return false
}
