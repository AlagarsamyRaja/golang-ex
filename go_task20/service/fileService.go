package service

import (
	"fileHandling/helperFunction"
	"fmt"
	"os"
)

func CreateFile() (string, error) {

	fmt.Println("Input the file name to be create :")
	fileName := helperFunction.ScanString()

	file, err := os.Create(fileName)
	if err != nil {
		return "", fmt.Errorf("ERROR::Failed to create the file: %s", err)

	}
	defer file.Close()

	fmt.Print("Input a sentence for the file :")
	sentence := helperFunction.ScanString()

	_, err = file.WriteString(sentence + "\n")
	if err != nil {
		return "", fmt.Errorf("ERROR::Failed to do write operation in the given file: %s", err)
	}

	return file.Name(), nil
}

func ReadFile() (string, string, error) {

	fmt.Println("Input the file name to be opened :")
	fileName := helperFunction.ScanString()

	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", "", fmt.Errorf("ERROR::Failed to read the given file check the file name: %s", err)
	}

	return fileName, string(content), nil
}

func AppendFile() (string, string, error) {

	fmt.Println("Input the file name to be opened :")
	fileName := helperFunction.ScanString()

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return "", "", fmt.Errorf("ERROR::Failed to open the file check the file name: %s", err)
	}

	// fmt.Println("Input the number of lines to be written :")
	// number := helperFunction.ScanString()

	// lineNumber, err := helperFunction.IsValidInput(number)
	// if err != nil {
	// 	return "", "", err
	// }

	fmt.Println("Enter the sentence to append in file and type end to exit")

	fmt.Println("The lines are:")
	for {
		lines := helperFunction.ScanString()

		if lines == "end" || lines == "End" {
			break
		}

		_, err := file.WriteString(lines + "\n")
		if err != nil {
			return "", "", fmt.Errorf("ERROR::Failed to append the given lines in the file :%s", err)
		}
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", "", fmt.Errorf("ERROR::Failed to read the given file check the file name :%s", err)
	}

	return fileName, string(content), err

}

// func CopyFile() (string, string, string, error) {

// 	fmt.Println("Input the source file name :")
// 	sourceFile := helperFunction.ScanString()

// 	fmt.Println("Input the new file name :")
// 	destinationFile := helperFunction.ScanString()

// 	source, err := os.Open(sourceFile)
// 	if err != nil {
// 		return "", "", "", fmt.Errorf("ERROR::Failed to read the given file check the file name :%s", err)
// 	}

// 	destination, err := os.Create(destinationFile)
// 	if err != nil {
// 		return "", "", "", fmt.Errorf("ERROR::Failed to create the fail :%s", err)
// 	}

// 	_, err = io.Copy(destination, source)
// 	if err != nil {
// 		return "", "", "", fmt.Errorf("ERROR::Failed to copy a file: %s", err)
// 	}

// 	content, err := os.ReadFile(destinationFile)
// 	if err != nil {
// 		return "", "", "", fmt.Errorf("ERROR::Failed to read the given file check the file name: %s", err)
// 	}

// 	return source.Name(), destination.Name(), string(content), err

// }

func CopyFile() (string, string, string, error) {

	fmt.Println("Input the source file name :")
	sourceFile := helperFunction.ScanString()

	fmt.Println("Input the new file name :")
	destinationFile := helperFunction.ScanString()

	source, err := os.ReadFile(sourceFile)
	if err != nil {
		return "", "", "", fmt.Errorf("ERROR::Failed to read the given file check the file name :%s", err)
	}

	// destination, err := os.Create(destinationFile)
	// if err != nil {
	// 	return "", "", "", fmt.Errorf("ERROR::Failed to create the fail :%s", err)
	// }

	err = os.WriteFile(destinationFile, source, 0666)
	if err != nil {
		return "", "", "", fmt.Errorf("ERROR::Failed to copy a file: %s", err)
	}

	content, err := os.ReadFile(destinationFile)
	if err != nil {
		return "", "", "", fmt.Errorf("ERROR::Failed to read the given file check the file name: %s", err)
	}

	return sourceFile, destinationFile, string(content), err

}

// func MergeFile() (string, error) {

// 	fmt.Println("Input the 1st file name :")
// 	firstFileName := helperFunction.ScanString()

// 	firstFile, err := os.Open(firstFileName)
// 	if err != nil {
// 		return "", fmt.Errorf("ERROR::Failed to read the given file check the file name: %s", err)
// 	}

// 	fmt.Println("Input the 2nd file name :")
// 	secondFileName := helperFunction.ScanString()

// 	secondFile, err := os.Open(secondFileName)
// 	if err != nil {
// 		return "", fmt.Errorf("ERROR::Failed to read the given file check the file name: %s", err)
// 	}

// 	fmt.Println("Input the new file name where to merge the above two files :")
// 	mergedFile := helperFunction.ScanString()

// 	mergedFiles, err := os.OpenFile(mergedFile, os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return "", fmt.Errorf("ERROR::Failed to create or write in the file: %s", err)
// 	}

// 	_, err = io.Copy(mergedFiles, firstFile)
// 	if err != nil {
// 		return "", fmt.Errorf("ERROR::Failed to copy a file: %s", err)
// 	}

// 	_, err = io.Copy(mergedFiles, secondFile)
// 	if err != nil {
// 		return "", fmt.Errorf("ERROR::Failed to copy a file: %s", err)
// 	}

// 	return mergedFile, err
// }

func MergeFile() (string, error) {

	fmt.Println("Input the 1st file name :")
	firstFileName := helperFunction.ScanString()

	firstFile, err := os.ReadFile(firstFileName)
	if err != nil {
		return "", fmt.Errorf("ERROR::Failed to read the given file check the file name: %s", err)
	}

	fmt.Println("Input the 2nd file name :")
	secondFileName := helperFunction.ScanString()

	secondFile, err := os.ReadFile(secondFileName)
	if err != nil {
		return "", fmt.Errorf("ERROR::Failed to read the given file check the file name: %s", err)
	}

	fmt.Println("Input the new file name where to merge the above two files :")
	mergedFile := helperFunction.ScanString()

	// mergedFiles, err := os.OpenFile(mergedFile, os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	return "", fmt.Errorf("ERROR::Failed to create or write in the file: %s", err)
	// }

	err = os.WriteFile(mergedFile, firstFile, 0666)
	if err != nil {
		return "", fmt.Errorf("ERROR::Failed to copy a file: %s", err)
	}

	err = os.WriteFile(mergedFile, secondFile, 0666)
	if err != nil {
		return "", fmt.Errorf("ERROR::Failed to copy a file: %s", err)
	}

	return mergedFile, err
}

func DeleteFile() (string, error) {

	fmt.Println("Input the file name to delete:")
	FileName := helperFunction.ScanString()

	err := os.Remove(FileName)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	return FileName, nil
}
