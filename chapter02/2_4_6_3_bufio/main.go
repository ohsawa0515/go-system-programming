package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("使用例\n")
	buffer.Flush()

	io.WriteString(buffer, "hogeほげ\n")
	buffer.Flush()
}
