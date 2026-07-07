package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src, _ := os.Open("input.txt")
	defer src.Close()

	dst, _ := os.Create("output.txt")
	defer dst.Close()

	// default buffer size 32KB
	n, err := io.Copy(dst, src)
	fmt.Printf("write %v byte", n)
	if err != nil {
		panic(err)
	}
}