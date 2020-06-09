//A method is a function with a special receiver argument.

package main

import "fmt"

type myInt int

type empData struct {
	empid   int
	empname string
}

func (n myInt) Sqr() myInt {
	return n * n
}

func (e1 empData) empInfo() empData {
	e1.empid = 005
	e1.empname = "Eva"
	return e1
}

func main() {
	f := (myInt)(5)
	fmt.Println(f.Sqr())

	e := empData{001, "Alexa"}
	fmt.Println(e.empInfo())

}
