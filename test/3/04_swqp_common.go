// 04_swqp_common.go
package main

import (
	"fmt"
)

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap: a = %v, b = %v\n", a, b)
}

func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Printf("main: swap: a = %v, b = %v\n", a, b)
}
