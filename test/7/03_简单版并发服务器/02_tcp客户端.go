package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err_dial := net.Dial("tcp", "127.0.0.1:8080")
	if err_dial != nil {
		fmt.Println("err_dial = ", err_dial)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024*2)
	buf = []byte("are you ok")
	_, err_write := conn.Write(buf)
	if err_write != nil {
		fmt.Println("err_write = ", err_write)
		return
	}

	n, err_read := conn.Read(buf)
	if err_read != nil {
		fmt.Println("err_read = ", err_read)
		return
	}
	fmt.Println(string(buf[:n]))
}
