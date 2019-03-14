package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "ajsjjsjsj")
	fmt.Println("slice = ", string(slice))

	str, _ := strconv.Atoi("m")
	fmt.Println("str = ", str)
}
