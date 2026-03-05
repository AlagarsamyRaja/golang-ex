package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func isValidInput(input string) (int, error) {
	input = strings.TrimSpace(input)
	number := strings.ReplaceAll(input, " ", "")
	num := strings.ReplaceAll(number, "\t", "")
	val, err := strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf("Invalid Input")
	}
	return val, nil
}

func displaySlice(slice []int) []int {
	fmt.Println("Diplay the given slice")
	fmt.Println(slice)
	return slice
}

func changeValueInTheSlice(slice []int) ([]int, error) {
	fmt.Println("Enter index:")
	scanner.Scan()
	value := scanner.Text()
	index, err := isValidInput(value)
	if err != nil {
		return slice, err
	}
	if index >= len(slice) {
		fmt.Println("ERROR:Given index is out of range")
	} else {
		fmt.Println("Enter value to update in slice")
		scanner.Scan()
		value := scanner.Text()
		val, err := isValidInput(value)
		if err != nil {
			return slice, err
		}
		fmt.Println("before updated")
		fmt.Println(slice)
		slice[index] = val
		fmt.Println("After updated")
		fmt.Println(slice)
	}
	return slice, nil
}

func copySlice(slice []int) {
	copySlice := slice[:]
	fmt.Println("After copySlice the slice")
	fmt.Println(copySlice)
}

func deleteSlice(slice []int) ([]int, error) {
	fmt.Println("Enter index to deleteSlice")
	scanner.Scan()
	value := scanner.Text()
	index, err := isValidInput(value)
	if err != nil {
		return slice, err
	}
	if index >= len(slice) {
		fmt.Println("Index out of range")
	} else {
		fmt.Println("before deleted")
		fmt.Println(slice)
		slice = append(slice[:index], slice[index+1:]...)
		fmt.Println("After deleted")
		fmt.Println(slice)
	}
	return slice, nil
}

func reverseSlice(slice []int) {
	reverse := make([]int, 0, len(slice))
	fmt.Println("Reverse the slice")
	for i := len(slice) - 1; i >= 0; i-- {
		reverse = append(reverse, slice[i])
	}
	fmt.Println(reverse)
}

func main() {
	fmt.Println("Enter size:")
	scanner.Scan()
	value := scanner.Text()
	size, err := isValidInput(value)
	if err != nil {
		fmt.Println(err)
		return
	}
	slice := make([]int, size)
	fmt.Println("Enter elements")
	for i := range size {
		scanner.Scan()
		value := scanner.Text()
		slice[i], err = isValidInput(value)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	for {
		fmt.Println("Enter 1 for display the given slice\nEnter 2 for change the value\nEnter 3 for copy the slice\nEnter 4 for delete the element\nEnter 5 for reverse the slice\nEnter 6 for Exit")
		scanner.Scan()
		value := scanner.Text()
		userChoice, err := isValidInput(value)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch userChoice {
		case 1:
			displaySlice(slice)

		case 2:
			slice, err = changeValueInTheSlice(slice)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			copySlice(slice)

		case 4:
			slice, err = deleteSlice(slice)
			if err != nil {
				fmt.Println(err)
			}
		case 5:
			reverseSlice(slice)

		case 6:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid")
		}
	}
}
