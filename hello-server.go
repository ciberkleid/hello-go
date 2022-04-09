package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

// global variables
var defaultMsg string = getEnv("HELLO_MSG", "sunshine")

func handler(w http.ResponseWriter, r *http.Request) {
    log.Print("Request received")
    msg := defaultMsg
    uriSegments := strings.Split(r.URL.Path, "/")
    if uriSegments[1] != "" {
        msg = uriSegments[1]
	log.Print("Request message: [%s]", msg)
    }
    fmt.Fprintf(w, "<h1>Hello %s!</h1>\n", msg)
}

func main() {
    log.Print("Server started. Default message: [%s]", defaultMsg)
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    mux.HandleFunc("/*", handler)
    log.Fatal(http.ListenAndServe(":8080", mux))
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
