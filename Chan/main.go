package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		sum := 10
		for i := 0; i < 10; i++ {
			sum += i
			fmt.Println("i", i)
			ch <- sum
		}
		// 谁生产谁关闭
		close(ch)
	}()

	// for v := range ch {
	// 	fmt.Println(v)
	// }

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}

}
