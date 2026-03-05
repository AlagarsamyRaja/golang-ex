package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPrime(i int) bool {
	var count int
	for j := 1; j <= i; j++ {
		if (i % j) == 0 {
			count++
		}
	}

	if count == 2 {
		return true
	}

	return false
}

func main() {
	var prime int
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter Number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number := strings.ReplaceAll(input, " ", "")
	
	num, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Enter number only")
		return
	}

	fmt.Println("Prime numbers are:")
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			prime++
			fmt.Print(i," ")
		}
	}
	fmt.Printf("\nNumber of prime numbers which are less than or equal to %d: %d",num,prime)
}
