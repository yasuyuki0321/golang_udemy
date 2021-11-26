package main

import (
	"fmt"
)

// 関数

// func 関数名(引数) 戻り値 {}
func Plus(x int, y int) int {
	return x + y
}

// 引数の型は、まとめることも可能
func Plus2(x, y int) int {
	return x + y
}

// 複数の戻り値がある場合
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値の変数を指定することも可能
// 変数を指定した場合は、returnの値はは不要
func Double(price int) (result int) {
	result = price * 2
	return
}

// 返り値のない関数
func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()
}
