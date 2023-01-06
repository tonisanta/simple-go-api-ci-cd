package main

import (
	"fmt"
	"net/http"
)

const version = 1

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":80", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world, version:%d!", version)
}
