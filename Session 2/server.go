package main

import (
	"fmt"
	"net"
	"time"
)


func main(){
	listener, err := net.Listen("tcp", "localhost:1234")

	if err != nil {
		return
	}

	defer listener.Close()


	for {
		conn, err := listener.Accept()

		if err != nil {
			return
		}

		conn.SetDeadline(time.Now().Add(5 * time.Second))
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	fmt.Println("Accepted")

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	if err != nil {
		return
	}

	fmt.Println("Received Data: " , string(buf))
}