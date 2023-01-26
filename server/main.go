package main

import (
	"log"
	"net/http"
)

const message = "Hello World "

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("server failed to start")
	}
}
