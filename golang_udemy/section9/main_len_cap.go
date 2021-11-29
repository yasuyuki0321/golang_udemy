package main

import "fmt"

// スライス
// append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	// スライス要素数が可変長
	// 配列はサイズの変更不可
	sl = append(sl, 300)
	fmt.Println(sl)

	// 複数の値を追加することも可能
	sl = append(sl, 400, 500)
	fmt.Println(sl)

	// makeスライスを作成する
	sl2 := make([]int, 5)
	fmt.Println(sl2)
	fmt.Printf("%T\n", sl2)

	/*
		var arr [3]int = [3]int{1, 2, 3}
		fmt.Printf("%T\n", arr)
	*/

	// 要素数を調べる
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

	// 第２引数が要素数、第3引数が容量
	sl3 := make([]int, 5, 10)
	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	// 要素数
	fmt.Println(len(sl3))
	// 容量
	// 事前にメモリを確保しておくことでパフォーマンスがよくなる
	// capは容量をオーバした場合2倍の値を確保する
	fmt.Println(cap(sl3))
}
