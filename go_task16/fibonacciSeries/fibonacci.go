package fibonacciSeries

import (
	"fmt"
	"goroutine/validation"
)

func GettingValue() (int, error) {

	val := validation.ScanString()
	value, err := validation.IsValidInput(val)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return value, nil
}

func CheckFibonacci(out chan int, num int, done chan bool) {
	a := -1
	b := 1

	for i := 1; i <= num; i++ {
		c := a + b
		out <- c
		a, b = b, c
	}
	close(out)

	done <- true
}
