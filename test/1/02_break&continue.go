// 02_break&continue.go
package main

import "fmt"

// import "time"

func main() {
	i := 0
	for {
		i++
		// time.Sleep(time.Second)
		if i > 100 {
			break
		} else {
			fmt.Println("i = ", i)
		}
	}
}
