package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "3.14daskdka 7. 5.7dlsldls "
	r := regexp.MustCompile(`\d+\.\d+`) //反引号代表原始字符串
	m := r.FindAllStringSubmatch(s, -1)
	// m := r.FindAllString(s, -1)
	fmt.Println("m = ", m)
}
