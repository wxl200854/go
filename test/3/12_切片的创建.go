package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))

	s = append(s, 100)
	fmt.Println("s = ", s)

	s1 := make([]int, 3, 5)
	fmt.Printf("len = %v, cap = %v\n", len(s1), cap(s1))

	s2 := make([]int, 3)
	fmt.Printf("len = %v, cap = %v\n", len(s2), cap(s2))
}
