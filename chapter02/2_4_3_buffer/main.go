package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer 例\n"))
	buffer.WriteString("buffer.WriteString 例\n")
	// buffer をポインタにして渡さないといけない
	io.WriteString(&buffer, "io.WriteString 例\n")
	fmt.Println(buffer.String())
}
