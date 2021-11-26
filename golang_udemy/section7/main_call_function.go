package main

import (
	"fmt"
)

// 関数
// 引数に関数を取る関数

// 引数に関数を取り、戻り値はなし
func CallFunction(f func()) {
	f()
}

func main() {
	// CallFunctionに無名関数を渡す
	CallFunction(func() {
		fmt.Println("I am a function")
	})
}
