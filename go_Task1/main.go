package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var name, address, education, company, designation string
	var mobileNumber, alternativeNumber, date, month, year int
	fmt.Print("Enter name:")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Enter your date of birth")
	fmt.Print("Enter date:")
	if _, err := fmt.Scanf("%d\n", &date); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter month:")
	if _, err := fmt.Scanf("%d\n", &month); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Enter year:")
	if _, err := fmt.Scanf("%d\n", &year); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d-%d-%d\n", date, month, year)
	fmt.Print("Enter address:")
	fmt.Scanf("%s\n", &address)
	fmt.Print("Enter education:")
	fmt.Scanf("%s\n", &education)
	fmt.Print("Enter current company:")
	fmt.Scanf("%s\n", &company)
	fmt.Print("Enter designation:")
	fmt.Scanf("%s\n", &designation)
	fmt.Print("Enter mobile number:")
	if _, err := fmt.Scanf("%d\n", &mobileNumber); err != nil {
		fmt.Println(err)
		return
	} else {
		var mobile = strconv.Itoa(mobileNumber)
		if len(mobile) != 10 {
			fmt.Println("mobile number must be 10 digit")
			return
		}
		fmt.Println(mobileNumber)
	}
	fmt.Print("Enter alternative phone:")
	if _, err := fmt.Scanln(&alternativeNumber); err != nil {
		fmt.Println(err)
		return
	} else {
		var alternative = strconv.Itoa(alternativeNumber)
		if len(alternative) != 10 {
			fmt.Println("mobile number must be 10 digit")
			return
		}
		fmt.Println(alternativeNumber)
	}
	fmt.Println("-------------------------")
	fmt.Printf("Name:%s\naddress:%s\neducation:%s\n", name, address, education)
	fmt.Println("-------------------------")
	integerInput := 5
	fmt.Printf("%d\n", integerInput)
	floatInput := 5.5
	fmt.Printf("%f\n", floatInput)
	floatingPointInput := 5.5
	fmt.Printf("%.2f\n", floatingPointInput)
	booleanInput := true
	fmt.Printf("%t\n", booleanInput)
	characterInput := 'a'
	fmt.Printf("%c\n", characterInput)
	stringInput := "alagar"
	fmt.Printf("%s\n", stringInput)
	arrayInput := [3]int{1, 2, 3}
	fmt.Printf("%v\n", arrayInput)
	sliceInput := []string{"apple", "orange", "mango"}
	fmt.Printf("%v\n", sliceInput)
	mapInput := map[int]string{1: "go", 2: "java", 3: "python"}
	fmt.Printf("%v\n", mapInput)
	var firstPerson Person
	var secondPerson Person
	firstPerson.name = "Alagar"
	firstPerson.age = 21
	firstPerson.job = "Go developer"
	firstPerson.salary = 10000

	secondPerson.name = "Deepak"
	secondPerson.age = 21
	secondPerson.job = "Go developer"
	secondPerson.salary = 10000

	fmt.Printf("%s\n", firstPerson.name)
	fmt.Printf("%d\n", firstPerson.age)
	fmt.Printf("%s\n", firstPerson.job)
	fmt.Printf("%d\n", firstPerson.salary)

	fmt.Printf("%s\n", secondPerson.name)
	fmt.Printf("%d\n", secondPerson.age)
	fmt.Printf("%s\n", secondPerson.job)
	fmt.Printf("%d\n", secondPerson.salary)
}
