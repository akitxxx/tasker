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
	mux.HandleFunc("/sign-up", handlers.SignUp)
	mux.HandleFunc("/sign-in", handlers.HandleAuth)
	mux.HandleFunc("/user", handlers.JwtMiddleware(handlers.GetUserList))
	mux.HandleFunc("/user/", handlers.JwtMiddleware(handlers.FindUserById))
	mux.HandleFunc("/task", handlers.JwtMiddleware(handlers.CreateTask))

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
		w.WriteHeader(http.StatusNotFound)
	}
	if r.Method == "GET" {
		w.Write([]byte("OK"))
	}
}
