package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err_dial := net.Dial("tcp", "127.0.0.1:8080")
	if err_dial != nil {
		fmt.Println("err_dial = ", err_dial)
		return
	}
	defer conn.Close()
	go func() {
		str := make([]byte, 1024)
		for {
			n, err_os := os.Stdin.Read(str)
			if err_os != nil {
				fmt.Println("err_os = ", err_os)
				return
			}
			conn.Write(str[:n])
		}
	}()
	buf := make([]byte, 1024)
	for {
		n, err_read := conn.Read(buf)
		if err_read != nil {
			fmt.Println("err_read = ", err_read)
			return
		}
		fmt.Println(string(buf[:n]))

	}
}
