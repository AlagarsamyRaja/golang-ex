package oddEven

import (
	"channel/validation"
	"fmt"
)

func SliceValues(slice []int, length int) {
	for i := range length {
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
