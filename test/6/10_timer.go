package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("现在是时间是：", time.Now())

	t := <-timer.C
	fmt.Println("t = ", t)
}
