package main

import "os"

func main() {
	os.Stdout.Write([]byte("os.Stdout 例\n"))
	os.Stderr.Write([]byte("os.Stderr 例\n"))
}
