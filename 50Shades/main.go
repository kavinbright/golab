package main

import "fmt"

/*
error:
./main.go:5:6: missing function body
./main.go:6:1: syntax error: unexpected semicolon or newline before {
*/
// func main()
// {
// 	fmt.Println("testing")
// }

// name := "zl" err: 简单声明只能在函数内使用
var name string
var age = 10

type info struct {
	addr string
}

func work() (string, error) {
	return "cz100", nil
}

func main() {
	// var a int
	// a = 10
	fmt.Printf("name: %s, age: %d", name, age)

	// 每次使用简短声明变量时需要有一个全新未声明的变量
	// str1 := "ss01"
	// str1 := "dd01" err

	str2 := "kk"
	str2, str3 := "kk2", "cc3"
	fmt.Printf("str1: %s, str3:%s", str2, str3)

	var infos info
	//  non-name info.addr on left side of :=
	//err: info.addr, err := work()
	var err error
	infos.addr, err = work()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("info.Addr: %s", infos.addr)
}
