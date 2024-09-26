package main

import (
	"bufio"
	"fmt"
	b1hangman "fr/alkaniel/hangman-cli/funcs"
	"os"
	"strings"
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
		fmt.Println(masked)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if strings.Contains(toFind, input) {
			masked = b1hangman.CheckEntry(input[0], toFind, masked)
			if !strings.Contains(masked, "_"){
				fmt.Println(masked + "\n")
				fmt.Println("Congrats")
				break
			}
		} else {
			Attempts--
		}
		if Attempts == 0 {
			b1hangman.PrintHangman(Attempts)
			fmt.Println("Game over")
			fmt.Println("====================")
			fmt.Printf("The word was : %v", toFind)
			break
		}
	}
}

