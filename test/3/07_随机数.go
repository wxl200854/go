package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //以当前时间作为种子
	for i := 0; i < 5; i++ {
		fmt.Println("rand = ", rand.Intn(100)) //限定在100以内的随机数
	}
}
