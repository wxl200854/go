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

//值传递
func (p Person) SetInfoValue(n string, a int, s byte) {
	p.name = n
	p.age = a
	p.sex = s
}

// 引用传递
func (p *Person) SetInfoPoint(n string, a int, s byte) {
	p.name = n
	p.age = a
	p.sex = s
}

func main() {
	p := Person{"mike", 18, 'm'}
	p.PrintInfo()
	p.SetInfoValue("John", 18, 'f')
	p.PrintInfo()
	(&p).SetInfoPoint("John", 22, 'm')
	p.PrintInfo()
}
