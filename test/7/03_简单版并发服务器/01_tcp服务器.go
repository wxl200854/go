package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println("addr success connect")
	buf := make([]byte, 1024*2)
	n, err_read := conn.Read(buf)
	if err_read != nil {
		fmt.Println("err_read = ", err_read)
		return
	}
	fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
}

func main() {
	ln, err_listen := net.Listen("tcp", "127.0.0.1:8080")
	if err_listen != nil {
		fmt.Println("err_listen = ", err_listen)
		return
	}
	defer ln.Close()
	for {
		conn, err_accept := ln.Accept()
		if err_accept != nil {
			fmt.Println("err_accept = ", err_accept)
			continue
		}

		go HandleConn(conn)
	}
}
