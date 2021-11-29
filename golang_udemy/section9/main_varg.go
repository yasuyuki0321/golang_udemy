package main

import "fmt"

// スライス
// 可変長引数

// ...型名 で可変長引数を取ることができる
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	fmt.Println(Sum())

	// スライスを渡すこともできる
	// ...はスライスを展開するという意味
	sl := []int{1, 2, 3, 4}
	fmt.Println(Sum(sl...))
}
