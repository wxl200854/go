package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}

func BubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	n := 10
	s := make([]int, n)
	InitData(s)
	fmt.Println("排序前： s = ", s)
	BubbleSort(s)
	fmt.Println("排序后： s = ", s)
}
