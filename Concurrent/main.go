package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// go tools.Sya("HELLO")
	// tools.Sya("CHINA")

	// -----------------chan-----------------------------------
	// arr1 := []int{7, 2, 8, -9, 4, 0}
	// ch := make(chan int)
	// go tools.Sum(arr1[:len(arr1)/2], ch)
	// go tools.Sum(arr1[len(arr1)/2:], ch)
	// x := <-ch
	// y := <-ch
	// fmt.Printf("x: %d, y: %d", x, y)

	fmt.Println("-----------------------------")
	fmt.Println("CPU:", runtime.NumCPU())
	// runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i <= 50; i++ {

			fmt.Println("A-i:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i <= 50; i++ {
			fmt.Println("B-i:", i)
		}
	}()

	wg.Wait()
}
