package main

import (
	"fmt"
)

//　型
// 文字列型

func main() {
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// 数字を文字列として扱う
	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	// 入植した文字列をそのまま出力する
	fmt.Println(`test
	test
		test`)

	// ダブルクオートを表示する
	fmt.Println("\"")
	fmt.Println(`"`)

	// stringはbyte型
	// 文字列はbyte配列のあつまり
	fmt.Println(string(s[0]))
}
