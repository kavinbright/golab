package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	a := [2]int{1, 2}
	fmt.Printf("a: %d", a[3])

	defer fmt.Println("YWS")
}
