package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  byte
}

type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s := Student{Person{"mike", 18, 'm'}, 1, "bj"}
	fmt.Println("s = ", s)
	fmt.Printf("s = %+v\n", s)

	var s2 Student
	s2.name = "tom"
	s2.addr = "sd"
	fmt.Println("s2 = ", s2)
	fmt.Printf("s2 = %+v\n", s2)
}
