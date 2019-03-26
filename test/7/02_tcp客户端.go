package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	buf := make([]byte, 1024)
	buf = []byte("jskkdskdkksadkkskdskdks")
	conn.Write(buf)
}
