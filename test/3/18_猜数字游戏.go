package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	var a int
	for {
		a = rand.Intn(10000)
		if a >= 1000 {
			break
		}
	}
	*p = a
}

func GetNum(s []int, num int) {
	s[0] = num / 1000
	s[1] = num % 1000 / 100
	s[2] = num % 100 / 10
	s[3] = num % 10
}

func OnGame(s []int) {
	var num int
	slice := make([]int, 4)
	for {
		for {
			fmt.Printf("请输入一个四位数：")
			fmt.Scan(&num)
			if num > 999 && num < 10000 {
				break
			} else {
				fmt.Println("您输入的不正确！！！")
			}
		}
		GetNum(slice, num)
		n := 0
		for i := 0; i < 4; i++ {
			if slice[i] > s[i] {
				fmt.Printf("您输入的第%d位大了\n", i+1)
			} else if slice[i] < s[i] {
				fmt.Printf("您输入的第%d位小了\n", i+1)
			} else {
				fmt.Printf("您输入的第%d位对了\n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("恭喜你，全部猜对了！！！")
			break
		}
	}

}

func main() {
	var randNum int

	CreateNum(&randNum)
	// fmt.Println("randNum = ", randNum)

	randSlice := make([]int, 4)
	GetNum(randSlice, randNum)
	// fmt.Println("randSlice = ", randSlice)

	OnGame(randSlice)
}
