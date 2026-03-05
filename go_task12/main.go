package main

import (
	"bufio"
	"fmt"
	"os"
	"pointer/check_length"
	"pointer/sort_array"
	"pointer/struct_with_ptr"
	"pointer/sum_of_two_num"
	"pointer/swapping"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	for {
		fmt.Println("Enter 1 for checking the length of the string\nEnter 2 for sorting the array\nEnter 3 for Sum of two values\nEnter 4 for using pointers in struct\nEnter 5 for swapping three numbers\nEnter 6 for exit ")
		fmt.Println("Enter your choice:")

		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		number := strings.ReplaceAll(input, " ", "")
		userChoice, _ := strconv.Atoi(number)

		switch userChoice {

		case 1:
			str := check_length.CheckLength()
			fmt.Printf("The length of the given string %v is %v\n", *str, len(*str))

		case 2:
			value, err := sort_array.SortArray()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("The elements in the array after sorting :")
			fmt.Println("Ascending Order")
			for i := range value {
				fmt.Printf("element - %d : %d\n", i+1, value[i])
			}

			fmt.Println("Descending Order")
			for i, j := len(value)-1, 1; i >= 0; i-- {
				fmt.Printf("element - %d : %d\n", j, value[i])
				j++
			}

		case 3:
			first, second, value, err := sum_of_two_num.Sum()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("The sum of %d and %d is %d\n", first, second, value)

		case 4:
			value, err := struct_with_ptr.Struct()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Employee Details:")
			fmt.Printf("Name: %v\n", value.Name)
			fmt.Printf("Date of birth: %v\n", value.Dateofbirth)
			fmt.Printf("Id: %v\n", value.Id)

		case 5:
			first, second, third, err := swapping.Swapping()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("The value after swapping are :")
			fmt.Println("Element 1:", *first)
			fmt.Println("Element 2:", *second)
			fmt.Println("Element 3:", *third)

		case 6:
			fmt.Println("End")
			return

		default:
			fmt.Println("Invalid")
		}
	}
}
