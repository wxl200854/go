package main

import (
	"errors"
	"fmt"
)

func Div(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("除数不能为0")
	} else {
		result = a / b
	}
	return
}

func main() {
	var a, b int
	fmt.Println("请输入两个数")
	fmt.Scan(&a, &b)
	result, err := Div(a, b)
	if err != nil {
		fmt.Printf("err = %v\n", err)
	} else {
		fmt.Printf("result = %v\n", result)
	}
}
