package main

import "fmt"

func MyFunc(args ...int) {
	fmt.Println("len(args) = ", len(args))
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}

func main() {
	MyFunc()
	MyFunc(1)
	MyFunc(1, 2, 3)
}
