package struct_with_ptr

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanString() string {
	scanner.Scan()
	str := scanner.Text()
	return str
}

func checkString(name string) string {
	strSlice := strings.Fields(name)
	spaceSeparated := strings.Join(strSlice, " ")
	return spaceSeparated
}

func checkDate(DateStr string) (time.Time, error) {
	birthDate, err := time.Parse("02-01-2006", DateStr)

	if err != nil {
		return time.Time{}, fmt.Errorf("Date must be in given format")
	}

	return birthDate, nil
}

func validateDate(date time.Time) (time.Time, error) {
	currentDate := time.Now()

	age := currentDate.Year() - date.Year()

	if age < 18 {
		return time.Time{}, fmt.Errorf("Age must be above 18")
	} else if age > 50 {
		return time.Time{}, fmt.Errorf("Age must be below 50")
	} else {
		return date, nil
	}

}

func isValidInput(input string) string {
	input = strings.TrimSpace(input)
	number := strings.ReplaceAll(input, " ", "")
	str := strings.ReplaceAll(number, "\t", "")
	return str
}

type Employee struct {
	Name        string
	Dateofbirth string
	Id          string
}

func Struct() (*Employee, error) {
	var emp Employee

	fmt.Println("Enter your name:")
	name := scanString()
	name = checkString(name)

	fmt.Println("Enter your date of birth:")

	birthDate := scanString()
	birth, err := checkDate(birthDate)
	if err != nil {
		return &emp, err
	}

	birth, err = validateDate(birth)
	if err != nil {
		return &emp, err
	}

	birthday := birth.Format("02-01-2006")

	fmt.Println("Enter your ID:")
	id := scanString()
	id = isValidInput(id)

	emp = Employee{Name: name, Dateofbirth: birthday, Id: id}

	structPointer := &emp

	return structPointer, nil
}
