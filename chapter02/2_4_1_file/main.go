package main

import "os"

func main() {
	file, err := os.Create("test.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	// バイト列を受けとるため、[]byte で変換している
	file.Write([]byte("os.File 使用例\n"))
}
