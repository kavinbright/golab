package tools

// Sum ...
func Sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	ch <- sum
}
