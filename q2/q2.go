package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	if text == "" {
		return 0, fmt.Errorf("texto vazio")
	}

	words := strings.Fields(text)
	numWords := len(words)

	if numWords == 0 {
		return 0, fmt.Errorf("texto vazio")
	}

	totalLetters := 0

	for _, word := range words {
		totalLetters += len(word)
	}

	average := float64(totalLetters) / float64(numWords)
	return average, nil
}
