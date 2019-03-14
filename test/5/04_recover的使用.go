package main

import (
	"fmt"
)

func testa() {
	fmt.Println("aaaaa")
}

func testb(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 100
}

func testc() {
	fmt.Println("ccccc")
}

func main() {
	testa()
	testb(20)
	testc()
}
