package main

import (
	"fmt"
)

func main() {

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	c := [5]int{1, 2, 3}
	fmt.Println("c = ", c)

	d := [5]int{2: 10, 4: 20}
	fmt.Println("d = ", d)
}
