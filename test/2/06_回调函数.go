// 05_回调函数.go
package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func cheng(a, b int) int {
	return a * b
}

func chu(a, b int) int {
	return a / b
}

type FuncType func(int, int) int

func Calc(a, b int, ftest FuncType) (result int) {
	fmt.Println("Calc")
	result = ftest(a, b)
	return result
}

func main() {
	result := Calc(10, 15, add)
	fmt.Println("Hello World!")
	fmt.Println("Calc = ", result)
}
