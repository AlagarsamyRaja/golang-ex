package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func scanInt() int {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return 0
	}

	input = strings.TrimSpace(input)

	integerArray, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return 0
	}
	return integerArray
}

func checkTarget(array [7]int, target int) {
	firstOcc := true
	found := false
	for i := range array {
		for j := i + 1; j < len(array); j++ {
			sum := array[i] + array[j]
			if sum == target {
				if firstOcc {
					fmt.Println("Shortest Path:", i, j, "-", array[i], array[j])
					firstOcc = false
					found = true
				} else {
					fmt.Println("Other Possibilities")
					fmt.Println(i, j, "-", array[i], array[j])
					found = true
				}
			}
		}
	}
	if !found {
		fmt.Println("Target value is not in array")
	}
}

func doubleValue(array [7]int) {
	for i := range array {
		if i+1 < len(array) {
			array[i+1] = array[i] * 2
		}
	}
	for i := range array {
		fmt.Printf("array_nums[%d] = %d\n", i, array[i])
	}
}

func sortAsc(array [6]int) {
	for i := range len(array) {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println("Ascending Order", array)
}

func sortDesc(array [6]int) {
	for i := range len(array) {
		for j := i + 1; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println("Descending Order", array)

}

func displayArray() {
	var integerArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	stringArray := [...]string{"red", "green", "orange", "yellow", "white"}
	fmt.Println("Integer Array", integerArray)
	fmt.Println("String Array", stringArray)
}

func target() {
	integerArray := [7]int{1, 2, 4, 5, 6, 7, 8}
	fmt.Println(integerArray)
	fmt.Println("Provide the target value")
	target := scanInt()
	checkTarget(integerArray, target)
}

func double() {
	integerArray := [7]int{}
	fmt.Println("Input the first element of the array:")
	firstNumber := scanInt()
	if firstNumber <= 0 {
		fmt.Println("Enter morethan 0")
	} else {
		integerArray[0] = firstNumber
		doubleValue(integerArray)
	}
}

func sort() {
	integerArray := [6]int{8, 3, 2, 1, 3, 2}
	fmt.Println(integerArray)
	sortAsc(integerArray)
	sortDesc(integerArray)
}

func main() {
	for {
		fmt.Println("Enter 1 for display the array\nEnter 2 for Matching the target in array\nEnter 3 for Double the value in array\nEnter 4 for sorting the array\nEnter 5 for exit ")
		fmt.Println("Enter your choice")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number := strings.ReplaceAll(input, " ", "")
		userChoice, _ := strconv.Atoi(number)
		switch userChoice {
		case 1:
			displayArray()

		case 2:
			target()

		case 3:
			double()

		case 4:
			sort()

		case 5:
			fmt.Println("End")
			return

		default:
			fmt.Println("Invalid")
		}
	}
}
