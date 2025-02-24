package main

import (
	"fmt"
	"net"
	EH "session7/handler"
)

func main(){
	tempServerAddr := `C:\Users\theo\AppData\Local\Temp\streaming_session_73372395151\28024.sock`
	client, err := net.Dial("unix", tempServerAddr)
	EH.ErrorHandler(err)	

	defer client.Close()
	
	msg := []byte("hello streaming")
	// _, err = client.Write(msg)
	// EH.ErrorHandler(err)

	for i := 0; i<3; i++{
		_, err = client.Write(msg)
		EH.ErrorHandler(err)
	}

	buf := make([]byte, 1024)
	for i := 0; i<3; i++ {
		n, err := client.Read(buf)
		EH.ErrorHandler(err)

		fmt.Println(string(buf[:n]))
	}
}