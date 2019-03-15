package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "annndclsldlsldlslldlannacccccanc"
	r := regexp.MustCompile(`a.c`) //反引号代表原始字符串
	m := r.FindAllStringSubmatch(s, -1)
	fmt.Println("m = ", m)
}
