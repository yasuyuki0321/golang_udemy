package main

import (
	"fmt"
)

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

// i5 := 500
var i6 int = 600

func main() {
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello GO"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名 := 値
	// 型を指定しなくても良い
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 同じ変数名で2回定義することはできない
	// i4 := 450
	// fmt.Println(i4)

	// i4 = "Hello"

	// 暗黙的な定義は、関数の外で定義することができない
	// fmt.Println(i5)

	// 明示的な定義であれば、関数の外で定義しても使える
	fmt.Println(i6)

	outer()

	// 変数は定義した関数内でのみアクセス可能
	// fmt.Println(s4)

	// 変数を定義して使用しないとエラーになる
	// var s5 string = "Not Use"
}
