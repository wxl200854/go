package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := a[2:5]
	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %v, cap = %v\n", len(s1), cap(s1))

	s1[1] = 666
	fmt.Println("s1 = ", s1)
	fmt.Println("a = ", a)

	s2 := s1[2:7]
	s2[2] = 777
	fmt.Println("s2 = ", s2)
	fmt.Println("a = ", a)
}
