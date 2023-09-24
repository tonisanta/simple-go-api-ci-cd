package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const Version = 2

func main() {
	http.HandleFunc("/hello", HelloHandler)
	log.Println(http.ListenAndServe(":80", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "Hello world, Version:%d, hostname:%s", Version, hostname)
}
