package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("子协程：i = ", i)
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("主协程：num = ", num)
	}
}
