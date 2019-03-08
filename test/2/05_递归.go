// 05_递归.go
package main

import (
	"fmt"
	// "time"
)

func test01(a int) int {
	if a == 1 {
		return a
	} else {
		return a + test01(a-1)
	}
}

func main() {
	fmt.Println(test01(100))
}
