package tools

import (
	"fmt"
	"time"
)

// Sya ...
func Sya(s string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%s \n", s)
	}
}
