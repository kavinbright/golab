package main

import "fmt"

func main() {
	sayHello()

	fmt.Println(f1())

	trtype()
}

func trtype() {
	type mytype int
	a := 12
	var b = mytype(a)
	fmt.Println("b---->:", b)
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func sayHello() {

	defer fmt.Println("test1: ", 1)
	defer fmt.Println("test1: ", 2)

	defer fmt.Println("test1: ", 3)

	defer fmt.Println("test1: ", 4)

}
