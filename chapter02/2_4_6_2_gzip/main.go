package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := gzip.NewWriter(file)
	defer writer.Close()
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer 使用例\n")
}
