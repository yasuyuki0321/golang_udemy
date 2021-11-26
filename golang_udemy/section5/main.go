package main

import (
	"fmt"
)

//定数

// 定数はグローバルに書くことが多い
// 変数の1文字目を大文字にすることで他のパッケージからも参照することができる
const Pi = 3.14

// 複数の定数を設定する
const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

// 値を省略すると前に設定した値が設定される
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 変数の場合、型の最大数を超えると定義ができない(定義の時点でエラー)
// 定数の場合は最大値はないので定義可能
//var Big_int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	// cannot assign to Pi
	// 定数は変更することができない
	//Pi = 3
	//fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	// 1 1 1 A A A
	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	// iotaは連続する整数を生成する
	fmt.Println(c0, c1, c2)
}
