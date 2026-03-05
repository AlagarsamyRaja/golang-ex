package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanString() string {
	scanner.Scan()
	str := scanner.Text()
	return str
}

func isValidInput(input string) (int, error) {
	input = strings.TrimSpace(input)
	number := strings.ReplaceAll(input, " ", "")
	num := strings.ReplaceAll(number, "\t", "")

	val, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid Input")
	}

	return int(val), nil

}

func findingLargestAndSmallestWord() {
	fmt.Print("Enter a sentence: ")
	sentence := scanString()
	fmt.Println("You entered:", sentence)

	words := strings.Split(sentence, " ")

	trimedWord := []string{}
	for i := range words {
		if len((words[i])) != 0 {
			trimedWord = append(trimedWord, words[i])
		}
	}

	min := trimedWord[0]
	max := trimedWord[0]
	for _, word := range trimedWord {
		if len(word) > len(max) {
			max = word
		}
		if len(word) < len(min) {
			min = word
		}
	}
	fmt.Println("Largest Word:", max)
	fmt.Println("Smallest Word:", min)

}

func findingRepeatedCharacter() {
	fmt.Print("Enter a sentence: ")

	sentence := scanString()
	sentence = strings.TrimSpace(sentence)
	words := strings.Split(sentence, " ")
	fmt.Println(words)

	wordMaps := make(map[string]string)
	for _, word := range words {
		max := 0
		repeat := ""
		fmt.Println(word)
		for i := 0; i < len(word); i++ {
			count := 0
			for j := i + 1; j < len(word); j++ {
				if word[i] == word[j] {
					count++
				}
			}
			if count > max {
				max = count
				repeat = string(word[i])
			}
		}
		wordMaps[word] = repeat
	}
	fmt.Println(wordMaps)

	for key, value := range wordMaps {
		fmt.Println(key, ":", value)
	}
}

func findingLargestNumber() {
	fmt.Println("Input the First integer:")

	value := scanString()

	first, err := isValidInput(value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Input the Second integer:")

	value = scanString()

	second, err := isValidInput(value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Input the Third integer:")

	value = scanString()
	third, err := isValidInput(value)
	if err != nil {
		fmt.Println(err)
		return
	}

	if first > second && first > third {
		fmt.Println("Maximum value of three integers:", first)
	} else if second > first && second > third {
		fmt.Println("Maximum value of three integers:", second)
	} else {
		fmt.Println("Maximum value of three integers:", third)
	}

}

func swapingTheNumbers() {
	fmt.Println("Input value for x:")

	value := scanString()

	first, err := isValidInput(value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Input value for y:")
	scanner.Scan()
	value = scanner.Text()

	second, err := isValidInput(value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Before swapping the value of x & y: %d & %d\n", first, second)

	first, second = second, first

	fmt.Printf("After swapping the value of x & y: %d & %d\n", first, second)
}

func findingMaximumAndMinimumNumberInArray() {
	fmt.Println("Enter size:")

	value := scanString()

	size, err := isValidInput(value)
	if err != nil {
		fmt.Println(err)
		return
	}

	array := make([]int, size)
	for i := range size {
		fmt.Printf("Element - %d :", i)
		scanner.Scan()
		value := scanner.Text()
		array[i], err = isValidInput(value)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	maximumValue := array[0]
	minimumValue := array[0]

	for i := range size {
		if array[i] > maximumValue {
			maximumValue = array[i]
		}
		if array[i] < minimumValue {
			minimumValue = array[i]
		}
	}

	fmt.Printf("Maximum element is: %d\n", maximumValue)
	fmt.Printf("Minimum element is: %d\n", minimumValue)
}

func main() {
	for {

		fmt.Println("Enter 1 for finding the largest and smallest word in a sentence\nEnter 2 for finding the repeated character for each word in the given sentence\nEnter 3 for swapping the two numbers\nEnter 4 for find maximum and minimum element in array\nEnter 5 for maximum of three numbers\nEnter 6 for exit")

		value := scanString()

		userChoice, err := isValidInput(value)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch userChoice {

		case 1:
			findingLargestAndSmallestWord()

		case 2:
			findingRepeatedCharacter()

		case 3:
			swapingTheNumbers()

		case 4:
			findingMaximumAndMinimumNumberInArray()

		case 5:
			findingLargestNumber()

		case 6:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid")
		}

	}

}
