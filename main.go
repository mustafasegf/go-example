package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// test
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello world\n")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(fmt.Sprint(":", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
