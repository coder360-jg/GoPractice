package main

import "fmt"

var (
	a bool   = true
	b string = "Hi"

	//int  int8  int16  int32  int64
	//uint uint8 uint16 uint32 uint64 uintptr
	c byte = 4 // same as uint8

	d rune = 77 //same as int32 - represents a Unicode code point
	//float32 float64
	e float64 = 827234.985943

	//complex64 complex128
	f complex64 = 5 + 8i
)

//Default values
var (
	aa bool
	bb string
	cc byte
	dd rune
	ee float64
	ff complex64
)

func main() {

	fmt.Printf("%T : %v\n", a, a)
	fmt.Printf("%T : %v\n", b, b)
	fmt.Printf("%T : %v\n", c, c)
	fmt.Printf("%T : %v\n", d, d)
	fmt.Printf("%T : %v\n", e, e)
	fmt.Printf("%T : %v\n", f, f)

	//Print Default values

	fmt.Printf("\n==Default values==\n\n")
	fmt.Printf("%T : %v\n", aa, aa)
	fmt.Printf("%T : %v\n", bb, bb)
	fmt.Printf("%T : %v\n", dd, dd)
	fmt.Printf("%T : %v\n", ee, ee)
	fmt.Printf("%T : %v\n", ff, ff)

}
