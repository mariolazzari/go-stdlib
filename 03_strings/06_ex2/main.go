package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "I'm not good at this"
	censoredWords := []string{"not"}
	censored := censorWords(sentence, censoredWords)
	fmt.Println((censored))

}

func censorWords(sentence string, censoredWords []string) string {
	censored := sentence

	for _, word := range censoredWords {
		if strings.Contains(sentence, word) {
			censored = strings.ReplaceAll(sentence, word, strings.Repeat("*", len(word)))
		}
	}

	return censored
}
