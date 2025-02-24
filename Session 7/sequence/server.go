package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	EH "session7/handler"
)

func main(){
	dir, err := ioutil.TempDir("", "sequence_session_7")
	EH.ErrorHandler(err)

	defer func(){
		err = os.RemoveAll(dir)
		EH.ErrorHandler(err)
	}()

	socket := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()))
	listener, err := net.Listen("unixpacket", socket)
	EH.ErrorHandler(err)

	fmt.Printf("Listening at : %s\n", listener.Addr())
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		EH.ErrorHandler(err)
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn){
	defer conn.Close()

	buf := make([]byte, 1024)
	
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client is done sending the message")
				return
			}
			EH.ErrorHandler(err)
		}

		_, err = conn.Write(buf[:n])
		EH.ErrorHandler(err)
	}
}