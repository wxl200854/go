package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(conn net.Conn, path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("文件发送完毕")
			} else {
				fmt.Println("f.Read err = ", err)
			}
			return
		}
		//给服务器发送内容
		conn.Write(buf[:n])
	}
}

func main() {
	//提示输入文件
	fmt.Println("请输入要传输的文件名：")
	var path string
	fmt.Scan(&path)

	//获取文件名 info.Name()
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}

	//主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err1 = ", err1)
		return
	}
	defer conn.Close()

	//给接收方先发送文件名
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn.Write err2 = ", err2)
		return
	}

	//接受对方的回复
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	if string(buf[:n]) == "ok" {
		fmt.Println("对方已经准备好接收")
		SendFile(conn, path)
	}
}
