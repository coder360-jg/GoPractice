package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	arr1 := arr[0:2]
	arr2 := arr[2:5]
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)

	arr1 = arr[0:3]
	arr1[2] = 7
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
