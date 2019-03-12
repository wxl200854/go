package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func test(s Student) {
	s.id = 0
	fmt.Println("test: s = ", s)
}

func test1(p *Student) {
	p.id = 0
	fmt.Println("test1: *p = ", *p)
}

func main() {
	var s Student = Student{1, "tom", 'm', 18, "Shandong"}
	fmt.Println("main: s = ", s)
	test(s)
	test1(&s)
}
