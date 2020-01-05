package main

import "fmt"

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	// インタフェースの型を持つ変数を宣言
	var talker Talker
	// インタフェースを満たす構造体のポインタは代入できる
	talker = &Greeter{name: "shuichi"}
	talker.Talk()
}
