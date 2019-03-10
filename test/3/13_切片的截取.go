// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	s := a[:6:6]
// 	fmt.Println("s = ", s)
// 	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))
// }

package main

import (
	"fmt"
)

func main() {
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	var slice = arr[1:3] //low = 1, hign = 3, mx = 5
	fmt.Println("arr=", arr)
	fmt.Println("slice=", slice)
	fmt.Println("slice len", len(slice))
	fmt.Println("slice cap", cap(slice))
}

// arr= [1 2 3 4 5]
// slice= [2 3]
// slice len 2
// slice cap 4
//[low:hign:max] 不指定max时，max默认为数组长度   此处为5
