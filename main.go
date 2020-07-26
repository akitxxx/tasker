package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lelouch99v/tasker/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", health)
	mux.HandleFunc("/auth", handlers.HandleAuth)
	mux.HandleFunc("/user", handlers.GetUserList)
	mux.HandleFunc("/user/", handlers.FindUserById)

	port := ":5010"
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Println("server is listening on port" + port + "...")
	if err := server.ListenAndServe(); err != nil {
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

func errorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)
}
