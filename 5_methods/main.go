//A method is a function with a special receiver argument.

package main

import "fmt"

type myInt int

func (n myInt) Sqr() myInt {
	return n * n
}

func main() {
	f := (myInt)(5)
	fmt.Println(f.Sqr())

}
