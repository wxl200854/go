// 10_defer_test.go
package main

import (
	"fmt"
)

func test(x int) {
	result := 10 / x
	fmt.Println("result = ", result)
}

func main() {
	defer fmt.Println("Hello World!")
	defer fmt.Println("你好！！")

	// test(5)
	defer test(0)

	defer fmt.Println("今天天气真好！！")

}
