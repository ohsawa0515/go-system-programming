package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder 例\n"))
	builder.WriteString("strings.Builder.WriteString 例\n")
	fmt.Println(builder.String())
}
