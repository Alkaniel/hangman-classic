package b1hangman

import (
	"bufio"
	"fmt"
	"os"
)

func PrintHangman(attempt int) {
	path := "ressources/hangman.txt"
	startReadLine := 1

	switch attempt {
	case 9:
		startReadLine = 1
	case 8:
		startReadLine = 9
	case 7:
		startReadLine = 17
	case 6:
		startReadLine = 25
	case 5:
		startReadLine = 33
	case 4:
		startReadLine = 41
	case 3:
		startReadLine = 49
	case 2:
		startReadLine = 57
	case 1:
		startReadLine = 65
	case 0:
		startReadLine = 73
	}

	lastReadLine := startReadLine + 7
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineRead := 1

	for scanner.Scan() {
		if lineRead >= startReadLine && lineRead <= lastReadLine {
			fmt.Println(scanner.Text())
		}
		lineRead++
	}
}
