package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/data"
	EH "main/handler"
	"net/http"
	"os"
	"strings"
	"log"
)

func main() {
	var name string
	var age int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input name : ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Input age : ")
	fmt.Scanf("%d\n", &age)

	person := data.Person{Name: name, Age: age}
	jsonData, err := json.Marshal(person)
	EH.ErrorHandler(err)

	// try log
	log.Printf("INFO - Client : JSON data created : %s\n", string(jsonData))

	reqBody := bytes.NewBuffer(jsonData)
	resp, err := http.Post("http://localhost:1234/sendFile", "application/json", reqBody)
	// EH.ErrorHandler(err)
	if(err != nil){
		log.Fatalf("ERROR - client : failed to send the request : %v", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	EH.ErrorHandler(err)
	fmt.Println("Server said : ", string(data))
	log.Println("INFO - Client : successfully received response")

}