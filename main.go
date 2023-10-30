package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// test
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	defer db.Close()

	if err := db.QueryRow(context.Background(), "SELECT 1").Scan(new(int)); err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request received from %s\n", req.RemoteAddr)
		fmt.Fprintf(w, "hello world!!! brilan\n")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("Server is running at :%s\n", port)

	err = http.ListenAndServe(fmt.Sprint(":", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
