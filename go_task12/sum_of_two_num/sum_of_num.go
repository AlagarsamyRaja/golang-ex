package sum_of_two_num

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

func add(first *int, second *int, sum *int) {
	*sum = *first + *second
}

func Sum() (int, int, int, error) {
	var firstNum, secondNum, result int

	fmt.Println("Input the first number:")
	firstNumber := scanString()

	firstNum, err := isValidInput(firstNumber)
	if err != nil {
		return 0, 0, 0, err
	}

	firstPointer := &firstNum

	fmt.Println("Input the second number:")
	secondNumber := scanString()

	secondNum, err = isValidInput(secondNumber)
	if err != nil {
		return 0, 0, 0, err
	}

	secondPointer := &secondNum

	add(firstPointer, secondPointer, &result)

	return firstNum, secondNum, result, err
}
