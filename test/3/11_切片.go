// 11_切片.go
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 5, 7, 9, 100}
	s := a[0:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) //切片的长度
	fmt.Println("cap(s) = ", cap(s)) //切片的容量

	fmt.Println("\n-----------------\n")

	s = a[1:2:4]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) //切片的长度
	fmt.Println("cap(s) = ", cap(s)) //切片的容量
}
