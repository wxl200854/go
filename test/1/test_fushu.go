// test_fushu.go
package main

import (
	"fmt"
)

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)
	// fmt.Printf("flag = %t\n", flag)
	// var t complex128
	// t = 1.22 + 3.14i
	// fmt.Println("t = ", t, "\n")
	// fmt.Println(real(t), imag(t))
	// fmt.Println("Hello World!")
	var a int
	fmt.Printf("请输入一个整数：")
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
