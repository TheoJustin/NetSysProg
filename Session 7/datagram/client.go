package main

import (
	"fmt"
	"net"
	"os"
	EH "session7/handler"
	"path/filepath"
)

func main(){
	tempDir := ``
	tempSocketFile := ``

	serverAddr, err := net.ResolveUnixAddr("unixgram", tempSocketFile)
	EH.ErrorHandler(err)

	clientSocket := filepath.Join(tempDir, fmt.Sprintf("c%d.sock", os.Getpid()))
	listener, err := net.ListenPacket("unixgram", clientSocket)
	EH.ErrorHandler(err)
	defer listener.Close()

	err = os.Chmod(clientSocket, os.ModeSocket|0622)
	EH.ErrorHandler(err)

	msg := []byte("hello datagram")
	for i := 0; i<3; i++{
		_, err := listener.WriteTo(msg, serverAddr)
		EH.ErrorHandler(err)
	}

	buf := make([]byte, 1024)
	for i := 0; i<3; i++{
		n, _, err := listener.ReadFrom(buf)
		EH.ErrorHandler(err)

		fmt.Println(string(buf[:n]))
	}
}