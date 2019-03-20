package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res := `{
 "company": "it",
 "subjects": [
  "Go",
  "C++",
  "Python",
  "Test"
 ],
 "isok": true,
 "price": 666.6666
}`

	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("m = ", m)

	var str string
	for key, value := m{
		
	}

}
