package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	// fmt.Println(list)
	// fmt.Println(len(list))
	if len(list) != 2 {
		fmt.Println("usage: xxx file")
		return
	}
	fileName := list[1]
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// fmt.Printf("name = %v, size = %v\n", info.Name(), info.Size())
	fmt.Println(info.Name())
}
