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

	val, err := strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf("Input must be in number")
	}

	return val, nil
}

type calculator interface {
	add() int
	sub() int
}

type values struct {
	firstNumber  int
	secondNumber int
}

func (s *values) add() int {
	return s.firstNumber + s.secondNumber
}

func (s *values) sub() int {
	return s.firstNumber - s.secondNumber
}

func main() {
	var cal calculator

	fmt.Println("Enter your choice:")
	fmt.Println("Enter 1 for Addition\nEnter 2 for Subtraction")
	userChoice := scanString()

	choice, err := isValidInput(userChoice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter the first value:")
	first := scanString()

	firstValue, err := isValidInput(first)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter the second value:")
	second := scanString()

	secondValue, err := isValidInput(second)
	if err != nil {
		fmt.Println(err)
		return
	}

	cal = &values{firstNumber: firstValue, secondNumber: secondValue}

	switch choice {

	case 1:
		sum := cal.add()
		fmt.Println("Addition:", sum)

	case 2:
		diff := cal.sub()
		fmt.Println("Subtraction:", diff)

	default:
		fmt.Println("Invalid choice")

	}
}
