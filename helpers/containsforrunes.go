package helpers

//Fonction contains qui marche avec une liste de runes.
func ContainsForHangman(letterNotMasked []string, letter string) bool {
	for _, r := range letterNotMasked {
		if r == letter {
			return true
		}
	}
	return false
}
