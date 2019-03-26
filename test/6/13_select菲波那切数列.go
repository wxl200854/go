package main

import (
	"fmt"
)

func fib(ch chan<- int, ok <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-ok:
			fmt.Println("结束")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	ok := make(chan bool)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("num = ", <-ch)
		}
		ok <- true
	}()

	fib(ch, ok)
}
