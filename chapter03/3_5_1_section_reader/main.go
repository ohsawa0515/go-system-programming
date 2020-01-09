package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	// 14バイトから7バイト取得する
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}
