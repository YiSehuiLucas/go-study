package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	src, _ := os.Open("input.txt")
	defer src.Close()

	dst, _ := os.Create("output.txt")
	defer dst.Close()

	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dst)
	defer writer.Flush()

	buf := make([]byte, 4096)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			writer.Write(buf[:n])
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("复制完成")
}
