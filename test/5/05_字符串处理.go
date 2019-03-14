package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"ac", "baby"}
	r := strings.Join(s, "+ ")
	fmt.Println(strings.Join(s, ", "))
	fmt.Println(r)
}
