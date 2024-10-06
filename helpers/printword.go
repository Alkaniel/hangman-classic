package helpers

import "fmt"

func PrintHangmanWord(word string) {
	for i, r := range word { // Parcours la liste de rune
		fmt.Print(string(r) + " ")
		if i < len(word)-1 { //Condition pour print seulement un espace
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}