// 09_json.go
package main

import (
	"encoding/json"
	"fmt"
)

type It struct {
	Company  string   `json:"company"` //此处的标签中间不能加空格，不然会出错
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	it := It{Company: "it", Subjects: []string{"Go", "C++", "Python", "Test"}, Isok: true, Price: 666.6666}
	// res, err := json.Marshal(it)
	res, err := json.MarshalIndent(it, "", " ")
	// res, err := json.Marshal(it)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("res = ", string(res))
}
