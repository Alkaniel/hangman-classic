package main

import (
	"bufio"
	"fmt"
	b1hangman "fr/alkaniel/hangman-cli/funcs"
	"os"
)

func main() {
	Attempts := 10
	fmt.Printf("Good Luck, you have %v attempts\n", Attempts)
	toFind := b1hangman.ChooseWordToFind(1)
	masked := b1hangman.MaskWord(toFind)

	for {
		if Attempts < 10 {
			b1hangman.PrintHangman(Attempts)
		}
		PrintRunesForHangman(masked)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose: ")
		input, _, _ := reader.ReadRune()
		if b1hangman.ContainsForRunes(toFind, input) {
			masked= b1hangman.CheckEntry(input, toFind, masked)
			if !b1hangman.ContainsForRunes(masked, '_'){
				PrintRunesForHangman(masked)
				fmt.Println("\nCongrats !")
				break
			}
		} else {
			Attempts--
		}
		if Attempts == 0 {
			b1hangman.PrintHangman(Attempts)
			fmt.Println("Game over")
			fmt.Println("====================")
			fmt.Print("The word was : ")
			PrintRunesForHangman(toFind)
			break
		}
	}
}

func PrintRunesForHangman(word []rune){
	for i, r := range word{
		fmt.Print(string(r) + " ")
		if i < len(word) - 1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
