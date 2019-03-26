package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if data, ok := <-ch; ok == true {
			fmt.Println("data = ", data)
		} else {
			break
		}
	}
}
