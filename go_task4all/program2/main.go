package main

import "fmt"

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

func main() {
	var size, first int
	fmt.Println("Enter the size")
	fmt.Scanf("%d\n", &size)
	arr := make([]int, size)
	fmt.Println("Input the first element of the array:")
	// fmt.Scanf("%d\n", &first)
	if _, err := fmt.Scan(&first); err != nil {
		fmt.Println(err)
		return
	}
	arr[0] = first
	doubleValue(arr)
	// for i:=0;i<size;i++{
	// 	if(i+1 < size){
	// 		arr[i+1]=arr[i]*2
	// 	}
	// }
	// for i:=0;i<size;i++{
	// 	fmt.Printf("array_nums[%d] = %d\n",i,arr[i])
	// }
}
