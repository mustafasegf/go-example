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

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request received from %s\n", req.RemoteAddr)

		poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Println("Unable to parse DATABASE_URL:", err)
			fmt.Fprintf(w, "Unable to parse DATABASE_URL: %s\n", err)
			return
		}

		db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err != nil {
			log.Println("Unable to create connection pool:", err)
			fmt.Fprintf(w, "Unable to create connection pool: %s\n", err)
			return
		}

		defer db.Close()

		if err := db.QueryRow(context.Background(), "SELECT 1").Scan(new(int)); err != nil {
			log.Println("Unable to connect to database:", err)
			fmt.Fprintf(w, "Unable to connect to database: %s\n", err)
			return
		}

		fmt.Fprintf(w, "hello world!!!\n")
	})

	http.HandleFunc("/db_url", func(w http.ResponseWriter, req *http.Request) {
		db_url := os.Getenv("DATABASE_URL")
		fmt.Fprintf(w, "DATABASE_URL: %s\n", db_url)
	})

	db_url := os.Getenv("DATABASE_URL")
	fmt.Println(db_url)

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
