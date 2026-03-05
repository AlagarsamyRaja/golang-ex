package helperFunction

import (
	"bufio"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func ScanString() string {

	scanner.Scan()
	str := scanner.Text()

	return str
}

// func IsValidInput(input string) (int, error) {

// 	input = strings.TrimSpace(input)
// 	number := strings.ReplaceAll(input, " ", "")
// 	num := strings.ReplaceAll(number, "\t", "")

// 	val, err := strconv.Atoi(num)
// 	if err != nil {
// 		return 0, fmt.Errorf("Input must be in number")
// 	}

// 	return val, nil
// }
