// 19_map.go
package main

import (
	"fmt"
)

func main() {
	m1 := map[int]string{1: "mike", 2: "john", 3: "lucy"}
	vale, ok := m1[1] //判断m[1]是否存在
	if ok {
		fmt.Printf("m1[1] = %v\n", vale)
	} else {
		fmt.Printf("不存在")
	}
	// fmt.Println("Hello World!")
}
