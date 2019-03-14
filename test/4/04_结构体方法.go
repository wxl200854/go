// 04_结构体方法.go
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  byte
}

func (p Person) PrintInfo() {
	fmt.Println("p = ", p)
}

func main() {
	p := Person{"mike", 18, 'm'}
	p.PrintInfo()
}
