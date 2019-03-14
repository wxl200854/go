package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("error: this is a test")
	fmt.Println("err = ", err)

	err1 := errors.New("test")
	fmt.Println("err1 = ", err1)
}
