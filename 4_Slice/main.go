package main

import "fmt"

func main() {

	arr := [6]int{1, 2, 3, 4, 5, 6} //array
	arr1 := arr[0:2]                //slice := []int{1, 2, 3, 4, 5, 6}
	arr2 := arr[2:5]
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)

	arr1 = arr[0:3]
	arr1[2] = 7
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)

	slc1 := arr[:] //by default size
	fmt.Println(slc1)

	slc1 = arr[:6]
	fmt.Println(slc1)

	slc1 = arr[0:]
	fmt.Println(slc1)

	slc1 = arr[0:6]
	fmt.Println(slc1)

	//The length of a slice is the number of elements it contains.
	//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	slc1 = arr[0:4]
	fmt.Printf("len %d ; capacity %d\n", len(slc1), cap(slc1))

	slc1 = arr[4:6]
	fmt.Printf("len %d ; capacity %d\n", len(slc1), cap(slc1))

	slc1 = arr[2:4]
	fmt.Printf("len %d ; capacity %d\n", len(slc1), cap(slc1))

}
