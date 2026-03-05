package check_length

import "fmt"

func CheckLength() *string {
	var input string

	fmt.Println("Input a string:")
	fmt.Scanf("%s\n", &input)

	pointerVariable := &input

	return pointerVariable
}
