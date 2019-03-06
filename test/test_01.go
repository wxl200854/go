package main

import "fmt"

func main() {
	var ch byte
	ch = 97
	fmt.Printf("%c, %d\n", ch, ch)

	var (
		a byte
		b byte
	)

	a, b = 'A', 'a'
	fmt.Printf("a = %d b = %d ", a, b)

}
