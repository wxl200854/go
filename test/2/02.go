package main

import "fmt"

func test(tmp ...int) {
	//myfuc(tmp...)
	myfuc2(tmp[1:3]...)
}

func myfuc(args ...int) {
	for _, data := range args {
		fmt.Println(data)
	}
}

func myfuc2(args ...int) {
	for _, data := range args {
		fmt.Println(data)
	}
}

func main() {
	test(1, 20, 3, 4)
}
