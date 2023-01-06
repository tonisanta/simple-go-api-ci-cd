package main

import (
	"fmt"
	"net/http"
)

const version = 2

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":80", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world, version:%d!", version)
}
