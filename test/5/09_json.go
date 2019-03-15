// 09_json.go
package main

import (
	"encoding/json"
	"fmt"
)

type It struct {
	Company  string
	Subjects []string
	Isok     bool
	Price    float64
}

func main() {
	it := It{"it", []string{"Go", "C++", "Python", "Test"}, true, 666.6666}
	// res, err := json.Marshal(it)
	res, err := json.MarshalIndent(it, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("res = ", string(res))
}
