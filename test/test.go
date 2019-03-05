package main

import "fmt"

func main() {
	var a int
	var b string
	fmt.Println("a = ", a)
	a = 10
	b = "welcome"
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 125 //自动推导类型
	c = 300
	fmt.Printf("a = %d, b = %s, c = %d", a, b, c)
}
