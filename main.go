package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	b1hangman "fr/alkaniel/hangman-cli/funcs"
	"fr/alkaniel/hangman-cli/helpers"
	"os"
	"strings"
)

var input string

type hangmanData struct {
	ToFind   string `json:"toFind"`
	Masked   string `json:"masked"`
	Attempts int `json:"attempts"`
}

var game = newGame()

/*
* Main est ma fonction coeur de mon code
Je fais beaucoup de fonction externe pour rester le plus propre possible
*/
func main() {
	if len(os.Args) > 1 && os.Args[1] == "--startWith" { //Permet de lancer une partie avec la précedente sauvegarde
		game.LoadGame()
		os.Remove("save.json")
		StartGameLoop(2) //Relance la partie précédemment lancée
	} else {
		game.UpdateToFind(b1hangman.ChooseWordToFind(1)) //Fonction qui choisi mon nombre (Il pourra être set plus tard avec la difficulté, c'est preset)
		game.UpdateMasked(b1hangman.MaskWord(game.ToFind))   //Fonction qui masque mon mot en laissant un n nombre de lettre. N = (longeur du mot)/2 - 1
		game.UpdateAttempts(10) //Fonction qui set le nombre d'essai
		StartGameLoop(1) //Lance une première partie
	}
}

func StartGameLoop(funct int) { //Fonction qui lance la boucle de jeu
	if funct == 1 { //
		fmt.Printf("Good Luck, you have %v attempts.\n", game.Attempts)
	} else if funct == 2 {
		fmt.Printf("Welcome back, you have %v attempts left.\n", game.Attempts)
	}
	for {
		if game.Attempts < 10 { //Permet de print le pendu que si on a moins de 10 essais
			b1hangman.PrintHangman(game.Attempts)
		}
		helpers.PrintHangmanWord(game.Masked) //Print le mot masqué
		input = inputHelper() //Fonction qui permet de choisir une lettre 
		if len(input) > 1 {
			if input == game.ToFind {
				fmt.Println("Congrats !")
				break
		    } else {
				game.Attempts -= 2
			}
		} else {
			if strings.Contains(game.ToFind, input) { 
				game.UpdateMasked(b1hangman.CheckEntry(input, game.ToFind, game.Masked)) 
				if !strings.Contains(game.Masked, "_") { 
					helpers.PrintHangmanWord(game.Masked) 
					fmt.Println("\nCongrats !")
					break
				}
			} else {
				game.Attempts--
			}
			if game.Attempts == 0 {
				b1hangman.PrintHangman(game.Attempts)
				fmt.Println("Game over")
				fmt.Println("====================")
				fmt.Print("The word was : ")
				helpers.PrintHangmanWord(game.ToFind)
				break
			}
		} 
	}
}

func inputHelper() string { //Fonction qui permet de choisir une lettre
	fmt.Print("Choose: ") 
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan() 
	input = strings.TrimSpace(reader.Text())
	if input == "STOP" { //Permet de stopper la partie et de sauvegarder
		game.SaveGame()
		fmt.Println("Game saved. Goodbye !")
		os.Exit(3)
	} else { 
		input = strings.ToLower(input)
		if len(input) == 0 {
			fmt.Println("Please enter a letter")
			return inputHelper()
		} else if len(input) > 1 {
			return input
		} else {
			input = string(input[0])
		}
	}
	return input
}

/** ============= Gestion des Saves ============= */

func newGame() *hangmanData { //Fonction qui permet de créer une nouvelle partie
	return &hangmanData{
		ToFind:   "debug",
		Masked:   "_____",
		Attempts: 10,
	}
}

func (game *hangmanData) LoadGame() { //Fonction qui permet de charger une partie sauvegardée
	data, err := os.ReadFile("save.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, game)
	if err != nil {
		panic(err)
	}
}

func (game *hangmanData) UpdateToFind(toFind string) { //Fonction qui permet de set le mot à trouver
	game.ToFind = toFind
}

func (game *hangmanData) UpdateMasked(masked string) { //Fonction qui permet de set le mot masqué
	game.Masked = masked
}

func (game *hangmanData) UpdateAttempts(attempts int) { //Fonction qui permet de set le nombre d'essais
	game.Attempts = attempts
}

func (game *hangmanData) SaveGame() { //Fonction qui permet de sauvegarder la partie
	data, err := json.Marshal(game)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("save.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
