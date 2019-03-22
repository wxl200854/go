package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()
	for i := 0; i < 2; i++ {
		runtime.Gosched() //出让时间片，等待其他协程完成再进行main协程
		fmt.Println("hello")
	}
}
