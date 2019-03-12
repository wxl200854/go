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
	var s Student = Student{1, "tom", 'm', 18, "Shandong"}
	fmt.Println("s = ", s)

	s1 := Student{id: 1, name: "mike"}
	fmt.Println("s1 = ", s1)
}
