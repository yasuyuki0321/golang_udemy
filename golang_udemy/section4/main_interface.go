package main

import (
	"fmt"
)

// interface & nil

func main() {
	// interface型を宣言
	// 初期値はnil
	var x interface{}
	fmt.Println(x)

	// 全ての型と互換性がある
	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)

	x = "A"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

	// invalid operation: x + 2 (mismatched types interface {} and int)
	// interfaceとint型で計算はできない
	// interfaceは演算で使用することができない
	// x = 2
	// fmt.Println(x + 2)
}
