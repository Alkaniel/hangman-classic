package main

import (
	"fmt"
	b1hangman "fr/alkaniel/hangman-cli/funcs"
)

func main() {
	Attempts := 3

	fmt.Printf("Good Luck, you have %v attempts", Attempts)

	b1hangman.PrintHangman(Attempts)
	toFind := b1hangman.ChooseWordToFind(1)
	fmt.Print(toFind)
}
