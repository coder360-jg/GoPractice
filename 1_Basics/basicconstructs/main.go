package main

import (
	"fmt"
)

func main() {

	var a, b, c int = 1, 2, 3 // 'var' and data tyoe
	fmt.Println(a, b, c)

	var d, e, f = 2.40, true, "hello" //only var
	fmt.Println(d, e, f)

	g, h, i := "go", 'c', 0xd // initialization. c-> 99 , 0xd->13
	fmt.Println(g, h, i)
}
