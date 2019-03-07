// 09_defer.go
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Hello World!")

	fmt.Println("你好！！")
}
