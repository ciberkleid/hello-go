package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("Env var: %s = %s \n", "environment", os.Getenv("environment"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, sunshine!\n")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
