package main

import (
	"fmt"
)

// 配列型

func main() {
	// 配列は要素数を変更できない
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)
	// [3]int ←が型名
	// [3]が型名に含まれるため、後から変更ができない
	// 配列の要素数を変更する場合にはスライス(可変長配列)を使う

	// 初期値を設定する
	// 文字列型の初期値は空文字
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	// 暗黙的な宣言
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素数を省略して宣言
	// 要素数を数えて ... に代入する
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// 値の取り出し
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])

	// 最後の要素の取り出し
	fmt.Println(arr2[len(arr2)-1])

	// 更新
	arr2[2] = "C"
	fmt.Println(arr2)

	// 同じint型でも要素数が違うとエラーになる
	// cannot use arr1 (type [3]int) as type [4]int in assignment
	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 配列の要素数を調べる
	fmt.Println(len(arr1))
	fmt.Printf("%T\n", arr1)
}
