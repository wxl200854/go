// 05_函数类型.go
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

type FuncType func(int, int) int

func main() {
	var result int
	result = add(1, 1)
	fmt.Println("result = ", result)

	var ftest FuncType
	ftest = add
	result = ftest(10, 20)
	fmt.Println("result1 = ", result)

	ftest = minus
	result = ftest(10, 5)
	fmt.Println("rsult3 = ", result)
}
