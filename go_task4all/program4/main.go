package main

import "fmt"

func checkTarget(arr []int, target int) {
	firstOcc := true
	found := false
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			if sum == target {
				if firstOcc {
					fmt.Println("Shortest Path:", i, j, "-", arr[i], arr[j])
					firstOcc = false
					found = true
				} else {
					fmt.Println("Other Possibilities")
					fmt.Println(i, j, "-", arr[i], arr[j])
					found = true
				}
			}
		}
	}
	if !found {
		fmt.Println("Invalid target")
	}
}

func doubleValue(arr []int) {
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) {
			arr[i+1] = arr[i] * 2
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("array_nums[%d] = %d\n", i, arr[i])
	}
}

func sortAsc(arr []int, size int) {
	if size == 1 {
		return
	}
	for i := 0; i < size-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	sortAsc(arr, size-1)
}

func sortDesc(arr []int, size int) {
	if size == 1 {
		return
	}
	for i := 0; i < size-1; i++ {
		if arr[i] < arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	sortDesc(arr, size-1)
}

func displayArray() {
	var arr1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := [...]string{"red", "green", "orange", "yellow", "white"}
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func target() {
	var size, target int
	fmt.Println("Enter the size")
	if _, err := fmt.Scan(&size); err != nil {
		fmt.Println(err)
		main()
	}
	arr := make([]int, size)
	fmt.Println("Enter elements")
	for i := 0; i < size; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(arr)
	fmt.Println("Provide the target:")
	if _, err := fmt.Scan(&target); err != nil {
		fmt.Println(err)
		main()
	}
	checkTarget(arr, target)
}

func double() {
	var size, first int
	fmt.Println("Enter the size")
	if _, err := fmt.Scan(&size); err != nil {
		fmt.Println(err)
		main()
	}
	arr := make([]int, size)
	fmt.Println("Input the first element of the array:")
	if _, err := fmt.Scan(&first); err != nil {
		fmt.Println(err)
		return
	}
	if first <= 0 {
		fmt.Println("Enter morethan 0")
	} else {
		arr[0] = first
		doubleValue(arr)
	}
}

func sort() {
	var size int
	fmt.Println("Enter the size")
	if _, err := fmt.Scan(&size); err != nil {
		fmt.Println(err)
		main()
	}
	arr := make([]int, size)
	fmt.Println("Enter elements")
	for i := 0; i < size; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(arr)
	sortAsc(arr, size)
	fmt.Println("Ascending Order", arr)
	sortDesc(arr, size)
	fmt.Println("Descending Order", arr)
}

func main() {
	var ch int
	for {
		fmt.Println("Enter 1 for display the array\nEnter 2 for Matching the target in array\nEnter 3 for Double the value in array\nEnter 4 for sorting the array\nEnter 5 for exit ")
		fmt.Println("Enter your choice")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			displayArray()

		case 2:
			target()

		case 3:
			double()

		case 4:
			sort()

		case 5:
			fmt.Println("End")
			return

		default:
			fmt.Println("Invalid")
		}

	}
}
