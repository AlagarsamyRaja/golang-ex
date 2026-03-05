package swapping

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

func swappingThreeNumbers(first *int, second *int, third *int) {
	fmt.Println("The value before swapping are :")
	fmt.Println("Element 1:", *first)
	fmt.Println("Element 2:", *second)
	fmt.Println("Element 3:", *third)

	*first, *second, *third = *third, *first, *second
}

func Swapping() (*int, *int, *int, error) {
	var firstNum, secondNum, thirdNum int
	var err error
	fmt.Println("Input the value of 1st element :")
	firstNumber := scanString()

	firstNum, err = isValidInput(firstNumber)
	if err != nil {
		return &firstNum, &secondNum, &thirdNum, err
	}

	fmt.Println("Input the value of 2nd element :")
	secondNumber := scanString()

	secondNum, err = isValidInput(secondNumber)
	if err != nil {
		return &firstNum, &secondNum, &thirdNum, err
	}

	fmt.Println("Input the value of 3rd element :")
	thirdNumber := scanString()

	thirdNum, err = isValidInput(thirdNumber)
	if err != nil {
		return &firstNum, &secondNum, &thirdNum, err
	}

	swappingThreeNumbers(&firstNum, &secondNum, &thirdNum)

	return &firstNum, &secondNum, &thirdNum, nil

}
