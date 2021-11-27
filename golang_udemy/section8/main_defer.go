package main

import (
	"fmt"
	"os"
)

// deffer
// 関数終了時に実行する処理を登録できる

func TestDefer() {
	defer fmt.Println("END") // 関数終了時に実行される
	fmt.Println("START")
}

func RunDefer() {
	// deferは後に登録されたものから実行される
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	TestDefer()

	/*
		// defferで複数の処理を登録する場合は無名関数で設定する
		defer func() {
			fmt.Println("1")
			fmt.Println("2")
			fmt.Println("3")
		}() // 無名関数の最後に()を付けると即時実行される
	*/

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("err")
	}
	defer file.Close()
	file.Write([]byte("Hello"))
}
