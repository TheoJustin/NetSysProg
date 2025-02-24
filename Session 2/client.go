package main

import (
	"net"
	"time"
)

func main(){
	dial, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		return
	}
	defer dial.Close()

	dial.SetWriteDeadline(time.Now().Add(3 * time.Second))

	message := "Hello, Server!"
	_, err = dial.Write([]byte(message))

	if err != nil {
		return
	}
}