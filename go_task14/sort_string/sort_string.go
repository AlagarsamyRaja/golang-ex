package sort_string

import (
	"fmt"
	"stringArray/validation"
)

func sortAsc(bytes []byte) string {

	for i := range bytes {
		for j := i + 1; j < len(bytes); j++ {
			if bytes[i] > bytes[j] {
				bytes[i], bytes[j] = bytes[j], bytes[i]
			}
		}
	}

	return string(bytes)
}

func SortString() string {

	fmt.Println("Input the string :")
	word := validation.ScanString()

	charWord := []byte(word)
	result := sortAsc(charWord)

	return result
}
