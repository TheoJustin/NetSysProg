package main

import (
	"fmt"
	"io/ioutil"
	EH "session7/handler"
	"net"
	"os"
	"path/filepath"
)

func main(){
	dir, err := ioutil.TempDir("", "datagram_session_7")
	EH.ErrorHandler(err)

	fmt.Printf("Temp dir : %s\n", dir)

	defer func(){
		err := os.RemoveAll(dir)
		EH.ErrorHandler(err)
	}()

	socket := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()))
	listener, err := net.ListenPacket("unixgram", socket)
	EH.ErrorHandler(err)
	fmt.Printf("Listening at : %s\n", listener.LocalAddr())
	defer listener.Close()

	err = os.Chmod(socket, os.ModeSocket|0622)
	EH.ErrorHandler(err)

	buf := make([]byte, 1024)
	for {
		n, client, err := listener.ReadFrom(buf)
		EH.ErrorHandler(err)

		_, err = listener.WriteTo(buf[:n], client)
		EH.ErrorHandler(err)
	}
}
