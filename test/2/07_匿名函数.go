// 07_匿名函数.go
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "Mike"

	// fmt.Println("Hello World!")

	f1 := func() {
		fmt.Printf("a = %v, b = %v\n", a, b)
	}

	f1()

	func() {
		fmt.Printf("a = %v, b = %v\n", a, b)
	}()

	m, n := func(x, y int) (max, min int) {
		if x > y {
			max = x
			min = y
		} else {
			max = y
			min = x
		}
		return
	}(10, 20)
	fmt.Printf("m = %v, n = %v\n", m, n)
}
