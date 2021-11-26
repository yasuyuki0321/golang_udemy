package main

import (
	"fmt"
)

// 関数
// クロージャーの実装

// 関数を返す関数
// 引数に文字列を受け取り、文字列を返す関数を返す関数
// 前回の処理の引数を返す
// 呼び出しのタイミングで初期化されない
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("Golang"))
}
