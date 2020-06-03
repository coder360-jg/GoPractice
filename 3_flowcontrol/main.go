package main

import (
	"fmt"
	"time"
)

func main() {
	forloop()
	fmt.Println()
	ifexample()
	fmt.Println()
	fmt.Println(ifexamplevariation(100))
	varscopeinifelse(88)

	fmt.Println(Sqrt(2))
	switchexample()
	fmt.Println()
	switchexample2()
	fmt.Println()
	deferexample()
	multipledefersexample()

}

//Forloop : Example C while is spelled for in Go
func forloop() {
	i := 0
	for i < 100 {
		fmt.Println(i)
		i = i + 10
	}
}

//Ifusage : Example for If use
func ifexample() {
	i := -10

	if i < 0 {
		fmt.Println((-1) * i)
	} else {
		fmt.Println(i)
	}

}

func ifexamplevariation(a int) int {
	if v := 100; a*a > v {
		return v // Scope of v limited till  here
	}
	return a
}

func varscopeinifelse(a int) {
	if v := 10; a*a < v {
		fmt.Println(a)
	} else {
		fmt.Println(v) // v is accessible in else part of for loop
	}
}

//Sqrt : Find sqaure root using newton's method z -= (z*z - x) / (2*z)
func Sqrt(x float64) float64 {
	z := 1.0
	z1 := 0.0
	counter := 0
	for z1 != z && counter < 10 {
		z1 = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z1, z)
		counter++
	}
	fmt.Println(counter)
	return z
}

func switchexample() {
	var str string
	switch fmt.Scan(&str); str { // switch can have string cases
	case "Hi":
		fmt.Println("Hi selected")
	case "Hello":
		fmt.Println("Hello selected")
	default:
		fmt.Println("Nothing selected")
	}
}

//  switch true : cleaner way to write if-then-else chains
func switchexample2() {

	tm := time.Now()
	fmt.Println(tm.Hour())
	switch { //switch true
	case tm.Hour() > 5 && tm.Hour() < 12:
		fmt.Println("its morning")
	case tm.Hour() >= 12 && tm.Hour() < 16:
		fmt.Println("its afternoon")
	case tm.Hour() >= 16 && tm.Hour() < 18:
		fmt.Println("its evening")
	default:
		fmt.Println("its night")
	}
}

//	defers the execution of a function until the surrounding function returns
func deferexample() {

	defer fmt.Println("Hello")
	fmt.Println("World")

}

// Deferred function calls are pushed onto a stack.
// deferred calls are executed in LIFO order.
func multipledefersexample() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
