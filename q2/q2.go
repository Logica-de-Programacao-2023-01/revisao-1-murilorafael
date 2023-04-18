package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {

	if text == "" {
		return 0, fmt.Errorf("texto vazio")
	}

	caracteres := []string{",", ".", "!", "?", ":", ";"}

	for i := 0; i < len(caracteres); i++ {
		text = strings.ReplaceAll(text, caracteres[i], "")
	}

	list := strings.Fields(text)

	sum := 0

	for i := 0; i < len(list); i++ {
		sum += len(list[i])
	}

	return float64(sum) / float64(len(list)), nil
}
