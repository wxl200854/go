package main

import (
	"fmt"
)

func main() {
	a := 10
	fmt.Printf("a = %d\n", a)
	fmt.Printf("&a = %v\n", &a)

	var p *int
	p = &a
	*p = 888
	fmt.Printf("a = %v\n", a)
}
