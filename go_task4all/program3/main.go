package main

import "fmt"

func chkTarget(arr []int, target int) {
	firstOcc := true
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			if sum == target {
				//fmt.Println(i,j,"-",arr[i],arr[j])
				if firstOcc {
					fmt.Println("Shortest Path:", i, j, "-", arr[i], arr[j])
					fmt.Println("Other Possibilities")
					firstOcc = false
				} else {
					fmt.Println(i, j, "-", arr[i], arr[j])
				}
			}else{
				fmt.Println("Given Target is invalid for the given array")
				return
			}

		}
	}
}

func main() {
	var size, target int
	fmt.Println("Enter the size")
	fmt.Scanf("%d\n", &size)
	arr := make([]int, size)
	fmt.Println("Enter elements")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)
	fmt.Println("Provide the target:")
	fmt.Scanf("\n%d", &target)
	chkTarget(arr, target)
}
