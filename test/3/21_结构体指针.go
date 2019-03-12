package main

import (
	"fmt"
)

func main() {
	type Student struct {
		id   int
		name string
		sex  byte
		age  int
		addr string
	}
	var s Student
	var p *Student
	p = &s
	p.id = 1
	p.age = 19
	p.addr = "bj"
	p.name = "mike"
	p.sex = 'm'

	fmt.Println("*p = ", *p)

	p2 := new(Student)
	p2.id = 1
	p2.age = 19
	p2.addr = "bj"
	p2.name = "mike"
	p2.sex = 'm'

	fmt.Println("*p2 = ", *p2)
}
