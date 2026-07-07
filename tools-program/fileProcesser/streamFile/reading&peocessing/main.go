package main

import (
	"fmt"
	"io"
	"os"
)

// 统计文件中字节数
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 4096)

	var total int64

	for {
		n, err := file.Read(buf)

		if n > 0 {
			// 处理当前读取的数据
			total += int64(n)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("文件大小：", total)
}