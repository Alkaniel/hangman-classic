package helpers

//Fonction ToUpper avec des runes.
func RuneToUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		r = r - 32
	}
	return r
}