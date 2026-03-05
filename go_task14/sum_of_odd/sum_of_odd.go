package sum_of_odd

import (
	"fmt"
	"stringArray/validation"
)

func oddSum(values int, sum *int) {

	if values%2 != 0 {
		*sum += values
	}

}

func OddSumValues() (int, error) {

	var sum int

	fmt.Println("Enter the first number:")
	firstNum := validation.ScanString()
	firstNumber, err := validation.IsValidInput(firstNum)
	if err != nil {
		return 0, err
	}

	oddSum(firstNumber, &sum)

	fmt.Println("Enter the second number:")
	secondNum := validation.ScanString()

	secondNumber, err := validation.IsValidInput(secondNum)
	if err != nil {
		return 0, err
	}

	oddSum(secondNumber, &sum)

	fmt.Println("Enter the third number:")
	thirdNum := validation.ScanString()

	thirdNumber, err := validation.IsValidInput(thirdNum)
	if err != nil {
		return 0, err
	}

	oddSum(thirdNumber, &sum)

	fmt.Println("Enter the fourth number:")
	fourthNum := validation.ScanString()

	fourthNumber, err := validation.IsValidInput(fourthNum)
	if err != nil {
		return 0, err
	}

	oddSum(fourthNumber, &sum)

	fmt.Println("Enter the fifth number:")
	fifthNum := validation.ScanString()

	fifthNumber, err := validation.IsValidInput(fifthNum)
	if err != nil {
		return 0, err
	}

	oddSum(fifthNumber, &sum)

	return sum, nil

}
