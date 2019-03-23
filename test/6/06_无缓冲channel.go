package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 4)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程：i = ", i)
			ch <- i
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("主协程：num = ", num)
	}
}
