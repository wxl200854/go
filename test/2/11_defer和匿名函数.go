// 11_defer和匿名函数.go
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Printf("内部： a = %v, b = %v\n", a, b)
	}(a, b)

	a = 111
	b = 222
	fmt.Printf("外部： a = %v, b = %v\n", a, b)
}

func main01() {
	a := 10
	b := 20

	defer func() {
		fmt.Printf("内部： a = %v, b = %v\n", a, b)
	}()

	a = 111
	b = 222
	fmt.Printf("外部： a = %v, b = %v\n", a, b)
}
