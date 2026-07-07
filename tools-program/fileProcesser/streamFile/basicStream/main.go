package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// stream copy file
	// open file
	src, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	// create dst file
	dst, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	// create buffer 4kB
	buf := make([]byte, 4*1024)

	for {
		// read a piece of data
		n, err := src.Read(buf)

		// write into file
		if n > 0 {
			_, writeErr := dst.Write(buf[:n])
			if writeErr != nil {
				panic(writeErr)
			}
		}

		// end judgement
		if err == io.EOF{
			break
		}
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("复制完成")
}
