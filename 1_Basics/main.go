package main

import (
	"1_Basics/basicoperation"
	"fmt"
)

func main() {
	str := "Go Corona Go" //Directly assigning value make variable of that type
	var count int = 100
	var a, b, ch int

	// Printing APIs
	fmt.Print("My First go code\n")
	fmt.Println(str)
	fmt.Printf("%d 0x%x %b\n", count, count, count)
	fmt.Println("\n === Learning how to access for loop in go ====")
	//Accessing Function from another package
	basicoperation.AccessforLoop()

	fmt.Printf("====  Choose operation ==== \n\n")
	fmt.Print("1. Add 2 numbers \n2. Subtract\n3. Multiply\n4. Divide\n5. Fibonacci\n Choice : ")

	//Reading using scan
	fmt.Scan(&ch)
	if ch != 5 {
		//Reading using scanf
		fmt.Scanf("%d \n %d", &a, &b)
	}

	// Using Switch case
	switch ch {
	case 1:
		fmt.Println(basicoperation.AddNumbers(a, b))
	case 2:
		fmt.Println(basicoperation.SubtractNumbers(a, b))
	case 3:
		fmt.Println(basicoperation.MultiplyNumbers(a, b))
	case 4:
		fmt.Println(basicoperation.DivideNumbers(a, b))
	case 5:
		basicoperation.PrintFibo()
	}

}
