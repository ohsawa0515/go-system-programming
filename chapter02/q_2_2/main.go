package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := io.MultiWriter(os.Stdout, file)
	csv := csv.NewWriter(w)
	csv.Write([]string{
		"csv.Write", "使用例", "example",
	})
	csv.Flush()

	csv.WriteAll([][]string{
		{
			"csv.WriteAll", "使用例", "example",
		},
		{
			"foo", "bar", "baz",
		},
	})
	csv.Flush()
}
