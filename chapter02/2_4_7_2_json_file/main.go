package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("example.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := io.MultiWriter(os.Stdout, file)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
		"使用":      "例",
	})
}
