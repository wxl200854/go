package main

import (
	"fmt"
	"time"
)

var ch = make(chan int) //定义全局变量只能用 var xxx = make(chan it)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)

	}
	fmt.Printf("\n")
}

func Person1() {
	Printer("Hello")
	ch <- 666
}

func Person2() {
	<-ch
	Printer("world!!!")
}

func main() {

	go Person1()
	go Person2()

	for {

	}
}
