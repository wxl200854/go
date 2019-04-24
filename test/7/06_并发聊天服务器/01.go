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

func Manager() { //转发消息
	OnlineMap = make(map[string]Client)
	for {
		msg := <-message
		for _, cli := range OnlineMap {
			cli.C <- msg
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + msg
	return
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	cliAddr := conn.RemoteAddr().String()
	cli := Client{make(chan string), cliAddr, cliAddr}
	OnlineMap[cliAddr] = cli
	go WriteMsgToClient(cli, conn)
	message <- MakeMsg(cli, "login")
	for {

	}
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

		go HandleConn(conn) //处理用户链接
	}
}
