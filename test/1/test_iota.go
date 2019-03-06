// test_iota.go
package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	fmt.Println("Hello World!")
}
