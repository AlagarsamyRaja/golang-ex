package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanString() string {
	scanner.Scan()
	str := scanner.Text()
	return str
}

func isValidInput(input string) (int, error) {
	input = strings.TrimSpace(input)
	number := strings.ReplaceAll(input, " ", "")
	num := strings.ReplaceAll(number, "\t", "")
	val, err := strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf("Input must be in number")
	}
	return val, nil
}

func findLastValue(number int) int{
	var lastDigit int
	if number<0{
		if number/10==0{
			lastDigit=number%10
			//fmt.Println(lastDigit)
		}else{
			lastDigit = -1*(number%10)
			//fmt.Println(lastDigit)
		}
	}else{
		lastDigit = number%10
	}
	return lastDigit
}

func sum(number int) int {
	if number == 0 {
		return 0
	}
	lastDigit := findLastValue(number)
	return lastDigit + sum(number/10)
}

func rev(s *string, index int) string {
	if index == 0 {
		return string((*s)[index])
	}
	return string((*s)[index]) + rev(s, index-1)
}


func main() {
	var input string

	fmt.Println("Enter number:")
	number := scanString()
	
	num, err := isValidInput(number)
	if err != nil {
		fmt.Println(err)
		return
	}

	sumResult := sum(num)
	fmt.Println("The sum of the digit is:",sumResult)

	fmt.Println("Enter a string:")
	fmt.Scanf("%s\n",&input)

	reversedString:=rev(&input, len(input)-1)
	fmt.Println("The reverse of the string is:",reversedString)
	fmt.Println("orignal String:",input)

}
