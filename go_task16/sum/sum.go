package sum

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

func Sum(numbers []int, sumChan chan int) {

	var sumNumbers int

	for _, n := range numbers {
		sumNumbers += n
	}

	sumChan <- sumNumbers

	close(sumChan)
}

func Printer(sumChan chan int, done chan bool) {

	value := <-sumChan

	fmt.Printf("Sum of the given numbers are: %d\n", value)

	done <- true
}
