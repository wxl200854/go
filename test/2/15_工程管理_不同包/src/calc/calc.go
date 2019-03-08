package calc

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func init() {
	fmt.Println("this is calc init")
}
