package main

import "fmt"

// スライス
// コピー

func main() {
	/*
		sl := []int{1, 2, 3, 4}
		sl2 := sl

		sl2[2] = 1000
		// [1 2 1000 4] 参照型なので元の値も変更されてしまう
		// 同じメモリアアドレスを指している
		fmt.Println(sl)

		var i int = 10
		i2 := i
		i2 = 100
		// 基本型の場合は別物として扱われる
		fmt.Println(i, i2)
	*/

	// 参照型: スライス、マップ、チャネル
	// 参照渡しなので、同じアドレスを参照する

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)
	// copy(コピー先, コピー元)
	n := copy(sl2, sl)
	// nはコピーに成功した数
	fmt.Println(n, sl2)
}
