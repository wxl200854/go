package main

import (
	"fmt"
	"io"
	"os"
)

func src(patha, pathb string) {
	f_src, err_src := os.Open(patha)
	if err_src != nil {
		fmt.Println("err = ", err_src)
		return
	}
	defer f_src.Close()
	f_dst, err_dst := os.Create(pathb)
	if err_dst != nil {
		fmt.Println("err = ", err_dst)
		return
	}
	defer f_dst.Close()
	buf := make([]byte, 1024*4)

	for {
		n, err := f_src.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		f_dst.Write(buf[:n])
	}

}

func main() {
	t := os.Args
	// fmt.Println(t[1], t[2])
	if len(t) != 3 {
		fmt.Println("usage : xxx srcfile dstfile")
		return
	}
	patha, pathb := t[1], t[2]
	if patha == pathb {
		fmt.Println("原文件名和目的文件名不能相同")
		return
	}
	src(patha, pathb)
}
