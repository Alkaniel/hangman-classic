package helpers

//Fonction contains qui marche avec une liste de runes.
func ContainsForRunes(letterNotMasked []rune, letter rune) bool {
	for _, r := range letterNotMasked {
		if r == letter {
			return true
		}
	}
	return false
}
