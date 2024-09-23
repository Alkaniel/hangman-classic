package main

import (
	"fmt"
	b1hangman "fr/alkaniel/hangman-cli/funcs"
)

func main() {
	Attempts := 10

	fmt.Printf("Good Luck, you have %v attempts", Attempts)

	b1hangman.PrintHangman(9)
}
