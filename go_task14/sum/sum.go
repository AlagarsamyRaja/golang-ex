package sum

import (
	"fmt"
	"stringArray/validation"
)

func add(first int, second int) int {

	sum := first + second

	return sum
}

func Sum() (int, error) {
	var firstNum, secondNum int

	fmt.Println("Input the first number:")
	firstNumber := validation.ScanString()

	firstNum, err := validation.IsValidInput(firstNumber)
	if err != nil {
		return 0, err
	}

	fmt.Println("Input the second number:")
	secondNumber := validation.ScanString()

	secondNum, err = validation.IsValidInput(secondNumber)
	if err != nil {
		return 0, err
	}

	result := add(firstNum, secondNum)

	return result, err
}
