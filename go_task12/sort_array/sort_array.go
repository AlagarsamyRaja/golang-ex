package sort_array

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanInt() (int, error) {
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)

	integerArray, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Input must be in number")
	}
	return integerArray, nil
}

func sortAsc(array *[5]int) {
	for i := range len(*array) {
		for j := i + 1; j < len(*array); j++ {
			if (*array)[i] > (*array)[j] {
				(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
			}
		}
	}
}

func SortArray() (*[5]int, error) {
	integerArray := [5]int{}

	fmt.Println("Input the number of elements to store in the array : 5")

	for i := range len(integerArray) {

		fmt.Printf("Element - %d:\n", i+1)
		value, err := scanInt()
		if err != nil {
			return nil, err
		}

		integerArray[i] = value
	}

	fmt.Println(integerArray)

	sortAsc(&integerArray)
	return &integerArray, nil
}
