package main

import "flag"

func main() {
	namePtr := flag.String("name", "username", "姓名")
	agePtr := flag.Int("age", 18, "年龄")
}
