package main

import (
	"fmt"
	"net/http"
	EH "session7/handler"
	"io"
)

func main(){
	resp, err := http.Get("https://localhost:1234/welcome")
	EH.ErrorHandler(err)

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	EH.ErrorHandler(err)

	fmt.Println("Server said : ", string(data))
}