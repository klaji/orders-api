package main

import (
	"fmt"
	"net/http"
)

func main(){
	server := &http.Server{
		Addr:":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("fail to listen",err)
	}
}

func basicHandler(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodGet {
		if r.URL.Path == "/foo" {

		}
	}
	if r.Method == http.MethodPost {
		if r.URL.Path == "/foo" {
			
		}
	}
	w.Write([]byte("Hello, World 5"))
}