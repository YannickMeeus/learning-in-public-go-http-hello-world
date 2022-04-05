package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/say-hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
	})
	port := "8080"
	log.Printf("Server started at port %v", port)
	http.ListenAndServe(":"+port, nil)
}
