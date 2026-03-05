package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scanInt() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num := scanner.Text()
	number := strings.ReplaceAll(num, " ", "")
	number = strings.ReplaceAll(number, "\t", "")
	input := strings.TrimSpace(number)

	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Enter numbers only")
		return 0, err
	} else {
		return value, nil
	}

}

type Stack struct {
	items []int
}

func (s *Stack) Push(num int) {
	s.items = append(s.items, num)
}

func (s *Stack) Peek() {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
	} else {
		top := s.items[len(s.items)-1]
		fmt.Printf("Top element in the stack is %d\n", top)
	}
}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
	} else {
		top := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		fmt.Println("popped Element:", top)
		fmt.Println("After Pop Opearion", s.items)
	}
}

func (s *Stack) display() {
	fmt.Println("Display the stack Elements")
	fmt.Println(s.items)
}

func (s *Stack) clear() {
	s.items = s.items[:0]
	fmt.Println("Stack elements are cleared")
}

func (s *Stack) checkLength() {
	fmt.Printf("The length of the stack is %d\n", len(s.items))
}

func (s *Stack) checkIsEmpty() {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Stack is not empty it has values")
}

func main() {
	var a Stack
	fmt.Println("Enter 1 for stack with default value\nEnter 2 for empty stack")
	opt, _ := scanInt()
	switch opt {
	case 1:
		a.items = append(a.items, 1, 2, 3, 4)
		fmt.Println("Stack with default values")
		fmt.Println(a.items)

	case 2:
		fmt.Println("Empty stack")

	default:
		fmt.Println("Invalid Option")
	}
	for {
		fmt.Println("Enter 1 for Push\nEnter 2 for Pop\nEnter 3 for Peek\nEnter 4 for Display\nEnter 5 for Clear\nEnter 6 for Checking the length of the stack\nEnter 7 for Checking the stack is empty or not\nEnter 8 for Exit")
		ch, _ := scanInt()
		switch ch {
		case 1:
			fmt.Println("Enter number to push the values")
			value, _ := scanInt()
			a.Push(value)
			fmt.Println("After push operation")
			fmt.Println(a.items)

		case 2:
			fmt.Println("After pop operation")
			a.Pop()

		case 3:
			fmt.Println("Peek Operation")
			a.Peek()

		case 4:
			a.display()

		case 5:
			if len(a.items) != 0 {
				a.clear()
				continue
			}
			fmt.Println("Stack is already empty")

		case 6:
			a.checkLength()

		case 7:
			a.checkIsEmpty()

		case 8:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid Option")
		}
	}
}
