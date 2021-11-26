package main

import (
	"fmt"
)

// 関数
// 関数を返す関数

func ReturnFunc() func() {
	return func() {
		fmt.Println("I an a function")
	}
}

func main() {
	// 関数が返される
	f := ReturnFunc()
	// 返された関数を実行
	f()
}
