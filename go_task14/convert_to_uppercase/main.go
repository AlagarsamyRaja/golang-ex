package convert_to_uppercase

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanString() string {

	scanner.Scan()
	str := scanner.Text()

	return str
}

func Uppercase() string {

	fmt.Println("Input a string in lowercase :")
	sentence := scanString()

	word := []byte(sentence)

	for i := range word {
		if word[i] >= 97 && word[i] <= 122 {
			word[i] = word[i] - 32
		}
	}

	return string(word)
}
