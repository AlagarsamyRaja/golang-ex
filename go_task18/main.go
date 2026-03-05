package main

import (
	"fmt"
	"sqlConnection/helperFunction"
	"sqlConnection/pkg"
	"sqlConnection/repository"

	_ "github.com/lib/pq"
)

func main() {

	pkg.InitDb()

	defer pkg.DB.Close()

	for {
		fmt.Println("Enter 1 for displaying all rows and columns\nEnter 2 for inserting rows in the table\nEnter 3 for Exit")
		fmt.Println("Enter choice:")

		choice := helperFunction.ScanString()
		userChoice, err := helperFunction.IsValidInput(choice)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch userChoice {

		case 1:

			rows, err := repository.DisplayRow()
			if err != nil {
				fmt.Println(err)
				continue
			}

			for i := range rows {
				fmt.Println(rows[i])
			}

		case 2:

			err := repository.InsertRow()
			if err != nil {
				fmt.Println(err)
				continue
			}

		case 3:

			fmt.Println("Exiting the program")
			return

		default:

			fmt.Println("Invalid option")

		}

	}
}
