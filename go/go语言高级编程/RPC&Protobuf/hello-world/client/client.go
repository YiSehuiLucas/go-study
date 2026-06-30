package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("dailing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "wochao", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}