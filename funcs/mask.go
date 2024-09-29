package b1hangman

import (
	"fr/alkaniel/hangman-cli/helpers"
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
			if helpers.ContainsForRunes(notmask, letter) { //Check si la rune du mot est la lune de la lettre.
				masked = append(masked, helpers.RuneToUpper(letter))
			} else {
				masked = append(masked, '_')
			}
		}
		return masked
	}
