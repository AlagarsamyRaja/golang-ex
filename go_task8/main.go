package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func scanString() string {
	scanner.Scan()
	str := scanner.Text()
	return str
}

func isValidInput(input string) (int, error) {
	input = strings.TrimSpace(input)
	number := strings.ReplaceAll(input, " ", "")
	num := strings.ReplaceAll(number, "\t", "")

	val, err := strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf("Key must be integer value")
	}

	return val, nil
}

func keyValue(maps map[int]string, num int) {

	for i := 1; i <= num; i++ {
		for {
			fmt.Println("Enter key for the map")
			keyString := scanString()
			key, err := isValidInput(keyString)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Enter value for the map")
			scanner.Scan()
			value := scanner.Text()
			maps[key] = value
			break
		}
	}

	fmt.Println("Display the given keys and value in the map")
	fmt.Println(maps)
}

func main() {
	var num int
	fmt.Println("Enter keys must be integer and values must be string for first map")
	detail := make(map[int]string)

	fmt.Println("Enter how many keys and values")
	fmt.Scanf("%d\n", &num)

	keyValue(detail, num)

	fmt.Println("Enter keys must be integer and values must be string for second map")
	fruit := make(map[int]string)

	fmt.Println("Enter how many keys and values")
	fmt.Scanf("%d\n", &num)

	keyValue(fruit, num)

	//Empty Map declaration
	emptyMap := map[string]string{}
	fmt.Printf("Empty map declaration:%v\n", emptyMap)

	//Map declaration using make function
	userDetails := make(map[int]string)
	fmt.Printf("Map is initialized using make:%v\n", userDetails)

	//adding the values in map
	fmt.Println("Adding values in the first map")

	for {
		fmt.Println("Enter key for the map:")

		keyString := scanString()

		key, err := isValidInput(keyString)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if _, exist := detail[key]; exist {
			fmt.Println("Key already exist")
			continue
		}

		fmt.Println("Enter value for the map")

		value := scanString()
		detail[key] = value

		fmt.Println("After added the values in first map")
		fmt.Printf("%v\n", detail)
		break
	}

	fmt.Println("---------------------------------------")

	//update
	fmt.Println("Updating values in the first map")
	for {
		fmt.Println("Enter key for update")
		scanner.Scan()
		keyString := scanner.Text()
		key, err := isValidInput(keyString)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if _, ok := detail[key]; !ok {
			fmt.Println("Key does not exist")
			continue
		}

		fmt.Println("Enter value to update")
		value := scanString()

		fmt.Println("Updating the values")
		detail[key] = value
		fmt.Printf("%v\n", detail)
		break
	}

	fmt.Println("---------------------------------------")

	for {
		//delete
		fmt.Println("Enter key to delete in first map")

		keyString := scanString()
		key, err := isValidInput(keyString)

		if err != nil {
			fmt.Println(err)
			continue
		}
		if _, exist := detail[key]; !exist {
			fmt.Println("Key does not exist")
			continue
		}

		fmt.Println("Deleted the key and values in first map")
		delete(detail, key)

		fmt.Printf("%v\n", detail)
		break
	}

	fmt.Println("---------------------------------------")

	//sorting
	fmt.Println("Sorting the first map")
	var keySlice []int
	for mapKey := range detail {
		keySlice = append(keySlice, mapKey)
	}

	for i := 0; i < len(keySlice); i++ {
		for j := i + 1; j < len(keySlice); j++ {
			if keySlice[i] < keySlice[j] {
				keySlice[i], keySlice[j] = keySlice[j], keySlice[i]
			}
		}
	}

	for _, k := range keySlice {
		fmt.Printf("%d : %s\n", k, detail[k])
	}

	fmt.Println("---------------------------------------")

	//merge
	mergeMap := make(map[int]string, len(detail)+len(fruit))
	maps.Copy(mergeMap, detail)

	maps.Copy(mergeMap, fruit)

	fmt.Println("After Merged two maps")
	fmt.Printf("%v\n", mergeMap)
}
