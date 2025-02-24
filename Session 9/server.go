package main

import (	
	"fmt"
	"net/http"
	EH "main/handler"
	"io"
	"main/data"
	"encoding/json"
	"log"
)

func sendFileHandler(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	EH.ErrorHandler(err)

	defer r.Body.Close()

	// ubah data json menjadi struct person
	var person data.Person
	err = json.Unmarshal(body, &person)
	EH.ErrorHandler(err)

	fmt.Printf("Name : %s\n", person.Name)
	fmt.Printf("Age : %d\n", person.Age)
	fmt.Fprintln(w, "Successfully received data")

	log.Println("INFO - sendFileHandler : successfully received data")
}

func middleware(method string, handlerFunc http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != method {
			fmt.Fprintf(w, "Method not allowed")

			// logging
			log.Println("ERROR - middleware : method not allowed")

			return
		}
		handlerFunc(w, r)
	}
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/sendFile", sendFileHandler)

	server := http.Server{
		Addr: "localhost:1234",
		Handler: mux,
	}

	log.Println("INFO - server : the server is running")

	err := server.ListenAndServe()
	EH.ErrorHandler(err)
}