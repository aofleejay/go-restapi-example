package main

import (
	"net/http"
)

func response(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello world"))
}

// func main() {
// 	http.HandleFunc("/", response)
// 	http.ListenAndServe(":8080", nil)
// }
