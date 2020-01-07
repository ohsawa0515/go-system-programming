package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	gw := gzip.NewWriter(w)
	defer gw.Close()
	gw.Header.Name = "test"
	mw := io.MultiWriter(gw, os.Stdout)
	encoder := json.NewEncoder(mw)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(source); err != nil {
		log.Print(err)
	}
	gw.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
