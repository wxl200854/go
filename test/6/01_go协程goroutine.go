package main

import (
	"fmt"
	"time"
)

func NewTask() {
	for {
		fmt.Println("new task")
		time.Sleep(time.Second)
	}
}

func main() {
	go NewTask()
	for {
		fmt.Println("main task")
		time.Sleep(time.Second)
	}
}
