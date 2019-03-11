// 16_copy.go
package main

import (
	"fmt"
)

func main() {
	srcSlice := []int{1, 2}
	dstSlice := []int{1, 1, 1}
	copy(dstSlice, srcSlice)
	fmt.Println("dstSlice = ", dstSlice)
}
