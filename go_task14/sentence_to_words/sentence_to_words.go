package sentence_to_words

import (
	"fmt"
	"stringArray/validation"
)

func SplitWords() []string {

	var word string
	var words []string

	fmt.Println("Input a string :")
	sentence := validation.ScanSentence()

	for i := 0; i < len(sentence); i++ {
		if sentence[i] != 32 {
			word += string(sentence[i])
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return words

}
