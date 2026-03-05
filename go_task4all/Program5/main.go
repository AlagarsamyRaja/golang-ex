package main

import "fmt"

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

func main() {
	var size int
	fmt.Println("Enter the size")
	fmt.Scanf("%d\n", &size)
	arr := make([]int, size)
	fmt.Println("Enter elements")
	for i := 0; i < size; i++ {
		// 	fmt.Scan(&arr[i])
		// }
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
