// 07_if实现类型断言.go
package main

import (
	"fmt"
)

type Student struct {
	name string
	id   int
}

func main() {
	slice := make([]interface{}, 3)
	slice[0] = 1
	slice[1] = "Welcome"
	slice[2] = Student{"mike", 666}

	for i, data := range slice {
		if value, ok := data.(int); ok == true {
			fmt.Printf("slice[%d] is int, value = %v\n", i, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("slice[%d] is string, value = %v\n", i, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("slice[%d] is Student, value = %v\n", i, value)
		} else {
			fmt.Println(".......")
		}
	}
}
