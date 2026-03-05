package main

import (
	"bufio"
	"channel/fibonacciSeries"
	"channel/oddEven"
	"channel/validation"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	for {
		fmt.Println("Enter 1 for Sending numbers into channels using go routines\nEnter 2 for find odd even using channels\nEnter 3 for fibonacci series using channels\nEnter 4 for Exit ")

		fmt.Printf("Enter your choice:")

		choice := validation.ScanString()

		userChoice, err := validation.IsValidInput(choice)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch userChoice {

		case 1:

			numberChannel := make(chan int, 10)

			go func(out chan int, capacity int) {

				for i := 1; i <= capacity; i++ {
					out <- i
				}

				close(out)

			}(numberChannel, cap(numberChannel))

			for i := range numberChannel {
				fmt.Println(i)
			}

		case 2:

			fmt.Println("Enter length for slice:")
			num := validation.ScanString()

			val, err := validation.IsValidInput(num)
			if err != nil {
				fmt.Println(err)
				continue
			}

			slice := make([]int, val)
			evenChannel := make(chan int)
			oddChannel := make(chan int)

			oddEven.SliceValues(slice, val)

			go oddEven.CheckOddEven(slice, evenChannel, oddChannel)

			for val := range slice {

				select {

				case val = <-oddChannel:
					fmt.Println("Odd:", val)

				case val = <-evenChannel:
					fmt.Println("Even:", val)
				}
			}

			val = <-oddChannel

		case 3:

			fibChannel := make(chan int, 7)

			go fibonacciSeries.CheckFibonacci(fibChannel, cap(fibChannel))

			for i := range fibChannel {
				fmt.Printf("%d\n", i)
			}

		case 4:

			fmt.Println("Exit")
			return

		default:

			fmt.Println("Invalid option")
		}
	}
}
