package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter 使用例")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
