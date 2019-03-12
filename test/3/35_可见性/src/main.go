package main

import (
	"fmt"
	"test"
)

func main() {
	test.Test()
	man := new(test.Man)
	man.Name = "kk"
	man.Age = 18
	fmt.Println("*man = ", *man)

}
