	package b1hangman

	import (
		"math/rand"
	)

	/** Fonction qui masque mon mot en laissant un n nombre de lettre. N = (longeur du mot)/2 - 1 */
	func MaskWord(toFind []rune) []rune {
		var masked []rune // Celle qui vas être return
		var notmask []rune //Liste de lettre à retirer
		n := len(toFind)/2 - 1 // Le nombre de lettre qui doit être reveal dans le mot
		for i:= 0; i < n; i++{   //La boucle s'arrête au nombre de 
			j := rand.Intn(len(toFind)-1) //Choisi un nombre aléatoire sur les mots
			notmask = append(notmask, toFind[j]) 
		}
		for _, letter := range toFind { //Liste pour masquer le mot
			if ContainsForRunes(notmask, letter) { //Check si la rune du mot est la lune de la lettre.
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
