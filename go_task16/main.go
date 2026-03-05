package main

import (
	"bufio"
	"fmt"
	"goroutine/fibonacciSeries"
	"goroutine/oddEven"
	"goroutine/sorting"
	"goroutine/sum"
	"goroutine/validation"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	for {
		fmt.Println("Enter 1 for Sum of the given numbers using go routines\nEnter 2 for Fibonacci series using channels\nEnter 3 for Find odd even using two go routines\nEnter 4 for sorting the array using go routines\nEnter 5 for Exit ")

		fmt.Printf("Enter your choice:")

		choice := validation.ScanString()

		userChoice, err := validation.IsValidInput(choice)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch userChoice {

		case 1:

			fmt.Println("Enter length for slice:")
			num := validation.ScanString()

			val, err := validation.IsValidInput(num)
			if err != nil {
				fmt.Println(err)
				continue
			}

			slice := make([]int, val)
			sumChannel := make(chan int)
			done := make(chan bool)

			sum.SliceValues(slice)

			go sum.Sum(slice, sumChannel)

			go sum.Printer(sumChannel, done)

			<-done

		case 2:

			fmt.Println("Enter number for fibonacci series:")

			value, err := fibonacciSeries.GettingValue()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fibChannel := make(chan int)
			done := make(chan bool)
			merge := make([]int, 0)

			go fibonacciSeries.CheckFibonacci(fibChannel, value, done)

			for range value {
				select {
				case num := <-fibChannel:
					merge = append(merge, num)

				case <-done:
					return
				}
			}

			fmt.Println("Fibonacci Series are:", merge)

		case 3:

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
			done := make(chan bool)

			oddEven.SliceValues(slice)

			go oddEven.CheckOddEven(slice, evenChannel, oddChannel)

			go oddEven.Printer(len(slice), evenChannel, oddChannel, done)

			<-done

		case 4:

			fmt.Println("Enter length for slice:")
			num := validation.ScanString()

			val, err := validation.IsValidInput(num)
			if err != nil {
				fmt.Println(err)
				continue
			}

			slice := make([]int, val)
			sortChannel := make(chan int)
			done := make(chan bool)

			sorting.SliceValues(slice)

			go sorting.SortArray(slice, sortChannel)

			go sorting.Printer(slice, sortChannel, done)

			<-done

		case 5:

			fmt.Println("Exit")
			return

		default:

			fmt.Println("Invalid option")
		}
	}
}
