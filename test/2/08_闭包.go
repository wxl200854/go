// 08_闭包.go
package main

import (
	"fmt"
)

func test() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := test()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
