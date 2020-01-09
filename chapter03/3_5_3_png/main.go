package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunk(file *os.File) []io.Reader {
	// チャンクを格納する配列
	var chunks []io.Reader

	// 最初の8バイト（シグネチャ）を飛ばす
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		if err := binary.Read(file, binary.BigEndian, &length); err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		// 次のチャンクに移動
		// 現在位置は長さを読み終わった箇所なので
		// チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動
		offset, _ = file.Seek(int64(length+8), 1)
	}

	return chunks
}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	chunks := readChunk(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
