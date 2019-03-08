// 12_获取命令行参数.go
package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args

	for i, data := range list {
		fmt.Printf("list[%v] = %v\n", i, data)
	}
	// fmt.Println("Hello World!")
}
