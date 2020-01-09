package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 32ビットのエンディアンのデータ（10000）
	data := []byte{0x0, 0x0, 0x27, 0x10}
	//data := []byte{0x10, 0x27, 0x0, 0x0}
	var i int32
	// エンディアンの変換
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
	fmt.Printf("raw data: %v", string(data))
}
