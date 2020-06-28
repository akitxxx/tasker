package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", health)
	http.HandleFunc("/hello", hello)

	port := "5010"
	fmt.Println("server is listening on port:" + port + "...")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println(r.URL.Path + " is not found.")
		errorHandler(w, r, http.StatusNotFound)
	}
	if r.Method == "GET" {
		w.Write([]byte("OK"))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("hello!"))
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)
}
