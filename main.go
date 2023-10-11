package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request received from %s\n", req.RemoteAddr)
		fmt.Fprintf(w, "hello world\n")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("Server is running at :%s\n", port)

	err := http.ListenAndServe(fmt.Sprint(":", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
