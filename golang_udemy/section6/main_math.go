package main

import (
	"fmt"
)

//　算術演算子

func main() {
	fmt.Println(1 + 2)
	// 文字列の結合
	fmt.Println("ABC" + "DE")

	fmt.Println(5 - 1)

	fmt.Println(5 * 4)
	fmt.Println(60 / 3)

	fmt.Println(9 % 4)

	n := 0
	n += 4 // n = n + 4
	fmt.Println(n)

	n++ // n = n + 1
	fmt.Println(n)

	n-- // n = n - 1
	fmt.Println(n)

	s := "ABC"
	s += "DEF" // s = "ABC" + "DEF"
	fmt.Println(s)
}
