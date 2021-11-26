package main

import (
	"fmt"
)

// 型
// 整数型

func main() {
	var i1 int = 100

	/*
			int8	8ビット符号付き整数	 -128 ～ 127
		    int16	16ビット符号付き整数 -32768 ～ 32767
		    int32	32ビット符号付き整数 -2147483648 ～ 2147483647
		    int64	64ビット符号付き整数 -9223372036854775808 ～ 9223372036854775807
	*/

	var i2 int64 = 200

	fmt.Println(i1 + 50)

	// 同じint型でもint(環境依存でint64)とint64では違うかたとみなされる
	// fmt.Println(i1 + i2)

	// %Tは型を表示する
	fmt.Printf("%T\n", i2)

	// 型変換
	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i1 + int(i2))

}
