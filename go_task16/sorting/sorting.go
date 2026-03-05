package sorting

import (
	"fmt"
	"goroutine/validation"
)

func SliceValues(slice []int) {

	for i := range slice {

		fmt.Printf("Enter value for index %d :", i)

		numbers := validation.ScanString()
		num, err := validation.IsValidInput(numbers)
		if err != nil {
			fmt.Println(err)
			return
		}

		slice[i] = num
	}

}

func SortArray(array []int, sort chan int) {

	for i := range array {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	for i := range array {
		sort <- array[i]
	}

	close(sort)
}

func Printer(array []int, sort chan int, done chan bool) {

	fmt.Println("After sorting an array")

	for value := range array {
		value = <-sort
		fmt.Println(value)
	}

	done <- true
}
