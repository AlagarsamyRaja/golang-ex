package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"stringArray/convert_to_uppercase"
	"stringArray/sentence_to_words"
	"stringArray/sort_string"
	"stringArray/sum"
	"stringArray/sum_of_odd"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	for {
		fmt.Println("Enter 1 for Sum of two values\nEnter 2 for convert a string to uppercase\nEnter 3 for Sum of odd numbers\nEnter 4 for sort the string in ascending order\nEnter 5 for split the sentence into words\nEnter 6 for exit ")
		fmt.Println("Enter your choice:")

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		number := strings.ReplaceAll(input, " ", "")
		userChoice, _ := strconv.Atoi(number)

		switch userChoice {

		case 1:

			result, err := sum.Sum()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("The sum of the above two integers = %d\n", result)

		case 2:

			result := convert_to_uppercase.Uppercase()
			fmt.Printf("Here is the above string in UPPERCASE : %s\n", result)

		case 3:

			result, err := sum_of_odd.OddSumValues()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Sum of all odd values: %d\n", result)

		case 4:

			result := sort_string.SortString()
			fmt.Printf("After sorting the string appears like : %s\n", result)

		case 5:

			result := sentence_to_words.SplitWords()
			for i := range result {
				fmt.Println(result[i])
			}

		case 6:

			fmt.Println("End")
			return

		default:

			fmt.Println("Invalid")
		}
	}
}
