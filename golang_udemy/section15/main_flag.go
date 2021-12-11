package main

import (
	"flag"
	"fmt"
)

// flag

func main() {
	// コマンドラインのオプション処理

	var (
		max int
		msg string
		x   bool
	)

	// IntVar　整数のオプション
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	// StringVar 文字列のオプション
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	// BoolVar bool型おｎオプション コマンドラインに与えられたらtrue なければfalse
	flag.BoolVar(&x, "x", false, "拡張オプション")

	flag.Parse()

	fmt.Println("処理数の最大値:", max)
	fmt.Println("処理メッセージ:", msg)
	fmt.Println("拡張オプション:", x)
}
