package main

import (
	"bufio"
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
	r_src := bufio.NewReader(f_src)

	f_dst, err_dst := os.Create(pathb)
	if err_dst != nil {
		fmt.Println("err = ", err_dst)
		return
	}
	defer f_dst.Close()
	for {
		buf, err_s := r_src.ReadBytes('\n')
		if err_s != nil {
			if err_s == io.EOF {
				break
			}
			fmt.Println("err = ", err_s)
		}
		f_dst.WriteString(string(buf))
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
