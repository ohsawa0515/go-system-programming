package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout, os.Stderr)
	io.WriteString(writer, "io.MultiWriter 使用例\n")
}
