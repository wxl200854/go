// 04.go
package main

import (
	"fmt"
)

func max(a, b int) (c int) {
	if a > b {
		c = a
	} else {
		c = b
	}
	return c
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	// max(a, b)
	fmt.Println("max = ", max(a, b))
}
