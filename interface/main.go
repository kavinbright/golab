package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int32
	adr  string
}

func main() {
	var person1 person
	person1.name = "CalenZhang"
	person1.age = 12
	person1.adr = "hubei"
	fmt.Println(person1)

	person2 := &person{}
	person2.name = "zhangliang"
	person2.adr = "MenJiang"
	person2.age = 18
	fmt.Println(*person2)

	person3 := person{}
	person3.name = "tangying"
	person3.adr = "Ch"
	person3.age = 18
	person31 := changeAdr(&person3)
	fmt.Println(person31)

	var person4 person
	person4.name = "侧大人"
	fmt.Println(person4)

	person5 := []person{}
	per5 := person{}
	per5.name = "大亮啊"
	person5 = append(person5, per5)
	fmt.Println(person5)

	person7 := person{}
	person7.name = "HaiCelan"
	fmt.Println(uperName(&person7))

	person8 := &person{
		name: "leDaren",
	}
	fmt.Println(uperName(person8))

}

func changeAdr(in *person) *person {
	in.adr = "MC.MZ"
	return in
}

func uperName(in *person) *person {
	in.name = strings.ToUpper(in.name)
	return in
}
