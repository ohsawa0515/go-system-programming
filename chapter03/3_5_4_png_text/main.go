package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
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
	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
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

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)

	// CRC を計算して追加
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())

	return &buffer
}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		log.Fatal()
	}
	defer newFile.Close()

	chunks := readChunk(file)

	// シグネチャ書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")

	// 先頭に必要なIHDRチャンクを書き込み
	io.Copy(newFile, chunks[0])

	// テキストチャンクを追加
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))

	// 残りのチャンクを追加
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	// 新しい画像のチャンクを表示
	newChunks := readChunk(newFile)
	for _, chunk := range newChunks {
		dumpChunk(chunk)
	}
}