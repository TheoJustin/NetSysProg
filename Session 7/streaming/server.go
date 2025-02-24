package main

import (
	"fmt"
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
	n, err := conn.Read(buf)
	EH.ErrorHandler(err)

	_, err = conn.Write(buf[:n])
	EH.ErrorHandler(err)
}