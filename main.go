package main

import (
	"bufio"
	"fmt"
	b1hangman "fr/alkaniel/hangman-cli/funcs"
	"fr/alkaniel/hangman-cli/helpers"
	"os"
)

/** Main est ma fonction coeur de mon code
Je fais beaucoup de fonction externe pour rester le plus propre possible */
func main() {
	Attempts := 10 //Le nombre de vie de base
	fmt.Printf("Good Luck, you have %v attempts\n", Attempts) //Phrase de début
	toFind := b1hangman.ChooseWordToFind(1) //Fonction qui choisi mon nombre (Il pourra être set plus tard avec la difficulté, c'est preset)
	masked := b1hangman.MaskWord(toFind) //Fonction qui masque mon mot en laissant un n nombre de lettre. N = (longeur du mot)/2 - 1

	for { // Début d'une boucle "infinie" dès qu'il y a "break", la boucle s'arrête
		if Attempts < 10 { 
			b1hangman.PrintHangman(Attempts) // Fonction qui print les différents états du pendu, si une vie est perdu.
		}
		PrintRunesForHangman(masked) //Fonction spéciale pour print un array de Rune
		reader := bufio.NewReader(os.Stdin) // Création d'un reader sur l'input
		fmt.Print("Choose: ")
		input, _, _ := reader.ReadRune() //Je set le reader pour qu'il lise une rune
		if helpers.ContainsForRunes(toFind, input) { //boucle si la rune est une lettre du hangman
			masked= b1hangman.CheckEntry(input, toFind, masked) //Fonction mettre a jour le mot avec la nouvelle lettre
			if !helpers.ContainsForRunes(masked, '_'){ //Checker si le mot est trouvé
				PrintRunesForHangman(masked)
				fmt.Println("\nCongrats !")
				break 
			}
		} else {
			Attempts--
		}
		if Attempts == 0 { // Si il n'y a plus de vie
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
	for i, r := range word{ // Parcours la liste de rune
		fmt.Print(string(r) + " ") 
		if i < len(word) - 1 { //Condition pour print seulement un espace 
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
