package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("error")
		return
	}
	defer f.Close()
	for i := 0; i < 10; i++ {
		buf := fmt.Sprintf("i = %v\n", i)
		f.WriteString(buf)
	}

}

func ReadFile(path string) {
	b := make([]byte, 1024*2)
	f, open_err := os.Open(path)
	if open_err != nil {
		fmt.Println("error")
		return
	}
	defer f.Close()
	n, read_err := f.Read(b)
	if read_err != nil && read_err != io.EOF {
		fmt.Println("read error")
		return
	}
	fmt.Println("b = ", string(b[:n]))
}

func ReadFileLine(path string) {
	f, open_err := os.Open(path)
	if open_err != nil {
		fmt.Println("error")
		return
	}
	defer f.Close()

	newf := bufio.NewReader(f)
	for {
		content, err := newf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error")
		}
		fmt.Printf("content = %v", string(content))
	}

}

func main() {
	path := "./tmp.txt"

	// WriteFile(path)
	// ReadFile(path)
	ReadFileLine(path)

}
