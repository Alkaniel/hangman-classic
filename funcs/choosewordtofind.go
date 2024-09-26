package b1hangman

import (
	"bufio"
	"math/rand"
	"os"
)

func ChooseWordToFind(difficulty int) string {
	path := "ressources/words.txt"

	switch difficulty {
	case 1:
		path = "ressources/words.txt"
	case 2:
		path = "ressources/words2.txt"
	case 3:
		path = "ressources/words3.txt"
	}

	toReturn := ""

	file, err := os.Open(path)
	if err != nil {
		toReturn = "Error opening file: open " + file.Name() + ": The system cannot find the file specified."
		return toReturn
	}
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	lineRead := 1
	lineCount := 0

	for scanner1.Scan() {
		lineCount++
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		toReturn = "Eroor opening file: open" + path + ": The system connot find the file specified"
		return toReturn
	}

	scanner2 := bufio.NewScanner(file)
	lineToRead := rand.Intn(lineCount) + 1

	lineRead = 1

	for scanner2.Scan() {
		if lineRead == lineToRead {
			toReturn = scanner2.Text()
			break
		}
		lineRead++
	}

	return toReturn
}
