package main

import (
	"fileHandling/helperFunction"
	"fileHandling/service"
	"fmt"
)

func main() {

	for {
		fmt.Println("Enter 1 for Creating the file\nEnter 2 for Reading the file\nEnter 3 for Append text into the file\nEnter 4 for Copy a file\nEnter 5 for Merge two files\nEnter 6 for Delete the file\nEnter 7 for Exit")

		fmt.Println("Enter choice:")

		choice := helperFunction.ScanString()

		userChoice, err := helperFunction.IsValidInput(choice)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch userChoice {

		case 1:

			fileName, err := service.CreateFile()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("The file", fileName, "create successfully...!!")

		case 2:

			fileName, content, err := service.ReadFile()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("The content of the", fileName, "is  :", content)

		case 3:

			fileName, content, err := service.AppendFile()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("The content of the", fileName, "is  :", content)

		case 4:

			fileName, copyFileName, content, err := service.CopyFile()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("The file", fileName, "copied successfully in the file", copyFileName)

			fmt.Println("If you read the new file you will see the content of the file :", content)

		case 5:

			fileName, err := service.MergeFile()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("The two files merged into", fileName, "file successfully..!!")

		case 6:

			fileName, err := service.DeleteFile()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("The file", fileName, "is deleted")

		case 7:

			fmt.Println("Exiting the program")
			return

		default:

			fmt.Println("Invalid Option")

		}
	}

}
