// 05_通过channel实现同步和数据交互.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	defer fmt.Println("主协程结束")
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("i = ", i)
			time.Sleep(time.Second)
		}
		ch <- "子协程完毕"
	}()
	str := <-ch
	fmt.Println("str = ", str)
}
