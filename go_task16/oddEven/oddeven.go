package oddEven

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

func CheckOddEven(nums []int, even chan int, odd chan int) {

	for _, n := range nums {
		if n%2 == 0 {
			even <- n
		} else {
			odd <- n
		}
	}

	close(even)

	close(odd)
}

func Printer(length int, even chan int, odd chan int, done chan bool) {
	for val := range length {

		select {
		case val = <-odd:
			fmt.Println("Odd:", val)

		case val = <-even:
			fmt.Println("Even:", val)
		}
	}
	done <- true
}
