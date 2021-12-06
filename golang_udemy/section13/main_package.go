package main

import (
	"fmt"
	"golang_udemy/section13/foo"
)

// スコープ

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	b = s
	{
		b := "BBB"
		fmt.Println(b)
	}
	fmt.Println((b))
	return b
}

func main() {
	fmt.Println(foo.Max)
	// f.Println(foo.min)

	fmt.Println(foo.ReturnMin())
	fmt.Println(appName())

	// 関数の中で定義された定数や変数は関数のなかからのみ呼び出し可能
	// fmt.Println(AppName, Version)

	fmt.Println(Do("AAA"))
}
