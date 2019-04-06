package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var OnlineMap map[string]Client

var message = make(chan string)

func Manager() {

}

func HandleConn(conn net.Conn) {

}

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer ln.Close()

	go Manager()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ln.Accept err = ", err)
			continue
		}

		go HandleConn(conn)
	}
}
