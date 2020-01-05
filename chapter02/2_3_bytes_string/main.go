package main

import "fmt"

func main() {
	byteArray := []byte("ASCII")
	fmt.Printf("%v\n", byteArray)

	str := string([]byte{0x41, 0x53, 0x43, 0x49, 0x49})
	fmt.Printf("%v\n", str)
}
