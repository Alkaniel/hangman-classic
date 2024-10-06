package b1hangman

import (
	"bufio"
	"math/rand"
	"os"
)
/** Fonction qui choisi mon mot (Il pourra être set plus tard avec la difficulté, c'est preset)*/
func ChooseWordToFind(difficulty int) string {
	path := "ressources/words.txt"

	switch difficulty { //Regarde le choix de la difficulté et renvoie vers le bon fichier
	case 1:
		path = "ressources/words.txt"
	case 2:
		path = "ressources/words2.txt"
	case 3:
		path = "ressources/words3.txt"
	}

	toReturn := ""

	file, err := os.Open(path) // Ouvre le fichier
 	if err != nil { // Gère si le fichier n'existe pas
		toReturn = "Error opening file: open " + file.Name() + ": The system cannot find the file specified."
		return string(toReturn)
	}
	defer file.Close()

	scanner1 := bufio.NewScanner(file) //Sert à lire le fichier
	lineRead := 1 //Variable de début de lecture
	lineCount := 0 // Compteur de ligne

	for scanner1.Scan() {
		lineCount++ //Compte le nombre de ligne
	}

	_, err = file.Seek(0, 0) //Reset le pointer pour éviter le OutOfBounds Error
	if err != nil {
		toReturn = "Eroor opening file: open" + path + ": The system connot find the file specified"
		return string(toReturn)
	}

	scanner2 := bufio.NewScanner(file) //Nouvelle lecture du fichier
	lineToRead := rand.Intn(lineCount) + 1 //Choix aléatoire du mot

	lineRead = 1 

	for scanner2.Scan() {
		if lineRead == lineToRead {//Check si la ligne choisi est égale à la ligne lue pour garder le mot.
			toReturn = scanner2.Text()
			break //sort du programme
		}
		lineRead++ // Sinon on descends.
	}
	return toReturn
}
