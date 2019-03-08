package main

import (
	// _ "calc"    //  _是引入calc包，但是不引用里面的函数，只是调用init函数
	"calc"
	"fmt"
)

func init() {
	fmt.Println("this is main init")
}

func main() {
	// result := calc.Add(10, 20)
	result := calc.Minus(10, 5)
	fmt.Println("result = ", result)
}
