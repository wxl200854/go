package main

import "fmt"

func main() {
	var a int
	var b float64
	var (
		m = 100
		n = 3.33
	)
	a, b = 10, 3.14

	var (
		c int
		d float64
	)
	c, d = 200, 20.22

	fmt.Println(a, b, c, d, m, n)
}
