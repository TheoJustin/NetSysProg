package main

import (
	"fmt"
	"net/http"
	EH "session7/handler"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome...")
}

func middleware(method string, handlerFunc http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != method {
			fmt.Println(w, "method not allowed")
			return
		}
		handlerFunc(w, r)
	}
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", middleware(http.MethodGet, welcomeHandler))

	server := http.Server{
		Addr: "localhost:1234",
		Handler: mux,
	
	}

	// certificate.pem buat enkripsi data
	// key.pem buat dekripsi data

	// server ngasi cert -> ke client
	// kalo cert aman dan terpercaya akan didekripsi oleh key
	// dan hasilnya akan dikirim oleh server akan didekripsi lagi dengan private key
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	EH.ErrorHandler(err)
}