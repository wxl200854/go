package main

import (
	"fmt"
)

type mystr string

type Person struct {
	name string
	age  int
	sex  byte
}

type Student struct {
	Person
	int
	mystr
}

func main() {
	s := Student{Person{"mike", 18, 'm'}, 1, "bj"}
	fmt.Println("s = ", s)
	fmt.Printf("s = %+v\n", s)

	var s2 Student
	s2.Person = Person{name: "tom"}
	s2.int = 666
	s2.mystr = "nihao"
	fmt.Println("s2 = ", s2)
	fmt.Printf("s2 = %+v\n", s2)
}
