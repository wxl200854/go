package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func ReceiveFile(conn net.Conn, path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err = ", err)
			}
			return
		}
		f.Write(buf[:n])
	}

}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer ln.Close()

	conn, err_a := ln.Accept()
	if err_a != nil {
		fmt.Println("err_a = ", err_a)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fileName := string(buf[:n])
	conn.Write([]byte("ok"))

	ReceiveFile(conn, fileName)
}
