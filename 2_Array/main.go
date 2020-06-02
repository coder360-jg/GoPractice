package main

import "fmt"

func main() {
	AccessArray()
}

//AccessArray Learnign how to access array
func AccessArray() {

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i * i
		fmt.Println(arr[i])
	}
}
