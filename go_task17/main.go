package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func ScanString() string {

	scanner.Scan()
	str := scanner.Text()

	return str
}

func IsValidInput(input string) (int, error) {

	input = strings.TrimSpace(input)
	number := strings.ReplaceAll(input, " ", "")
	num := strings.ReplaceAll(number, "\t", "")

	val, err := strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf("Input must be in number")
	}

	return val, nil
}

func main() {

	fmt.Println("Input number of lines:")
	rows := ScanString()
	row, err := IsValidInput(rows)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Number of characters in a line:")
	columns := ScanString()
	cols, err := IsValidInput(columns)
	if err != nil {
		fmt.Println(err)
		return
	}

	// num:=1
	// for range row{
	// 	for range cols{
	// 		fmt.Printf("%d ",num)
	// 		num++
	// 	}
	// 	fmt.Println()
	// }

	array := make([][]int, row)
	for i := range array {
		array[i] = make([]int, cols)
	}

	num := 1
	for i := range row {
		for j := range cols {
			array[i][j] = num
			num += 1
		}
	}

	for i := range row {
		for j := range cols {
			fmt.Printf("%d ", array[i][j])
		}
		fmt.Println()
	}

}
