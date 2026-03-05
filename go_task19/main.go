package main

import (
	"bufio"
	"fmt"
	"logs/logs"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func ScanString() string {

	scanner.Scan()
	str := scanner.Text()

	return str
}

func main() {

	fmt.Println("Enter the text:")
	text:=ScanString()

	logs.Logs(text)

}