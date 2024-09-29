package b1hangman

/** CheckEntry est un checker qui regarde si la proposition correspond à l'une des lettre du mot à chercher
Ce qui renvoi le mot masqué avec les lettres en démasqué.*/
func CheckEntry(input rune, toFind []rune, masked []rune) []rune {
	for i, char := range toFind {
		if char == input {
			masked[i] = RuneToUpper(input)
		}
	}
	return masked
}