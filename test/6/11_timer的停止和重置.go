package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("时间到")
	}()
	//timer.Stop()
	timer.Reset(1 * time.Second)
	for {

	}
}
