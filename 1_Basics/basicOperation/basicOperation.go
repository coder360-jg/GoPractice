package basicOperation

import (
	"fmt"
)

func init() {

}

//AccessforLoop to print elements using loop
func AccessforLoop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}
}

//PrintFibo to calculate fibonacci  series
func PrintFibo() {
	fmt.Println("\n === Calling  fibonacci series ====")
	for i := 1; i < 10; i++ {
		fmt.Printf("%d  %d\n", i, fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return (fibonacci(n-1) + fibonacci(n-2))
}

//AddNumbers : Adds 2 nos
func AddNumbers(a, b int) int {
	return a + b
}

//SubtractNumbers : Sub two nos.
func SubtractNumbers(a, b int) int {
	return a - b
}

//MultiplyNumbers : Multiply 2 nos.
func MultiplyNumbers(a, b int) int {
	return a * b
}

//DivideNumbers : Divide 2 numbers
func DivideNumbers(a, b int) int {
	if b == 0 {
		fmt.Println("Undefined")
		return 0
	} else {
		return a / b
	}
}
