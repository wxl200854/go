// 09_指针数组做函数参数.go
package main

import (
	"fmt"
)

func modify(a [5]int) {
	a[0] = 666
	fmt.Println("a = ", a)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a)
	fmt.Println("main: a = ", a)
}
