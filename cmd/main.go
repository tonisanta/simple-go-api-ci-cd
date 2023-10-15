package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const Version = 7

func main() {
	log.Printf("Hello from version: %d", Version)
	http.HandleFunc("/hello", HelloHandler)
	log.Println(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println(err)
	}

	defaultMsg := "Hello world"
	value, set := os.LookupEnv("message")
	if !set {
		value = defaultMsg
	}

	fmt.Fprintf(w, "%s, Version:%d, hostname:%s", value, Version, hostname)
}
