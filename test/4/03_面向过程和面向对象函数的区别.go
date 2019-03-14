// 03_面向过程和面向对象函数的区别.go
package main

import (
	"fmt"
)

type long int

func Add01(a, b int) (result int) {
	result = a + b
	return
}

func (tmp long) Add02(other long) (result long) {
	result = tmp + other
	return
}

func main() {
	var result int
	result = Add01(10, 20)
	fmt.Println("result = ", result)

	var a long = 5
	r := a.Add02(10)
	fmt.Println("r = ", r)
}
