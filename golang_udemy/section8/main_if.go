package main

import "fmt"

// if
// 条件分岐

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	// 簡易分付きif文
	// 条件式の前に簡易分を設定できる
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	// 条件付きif文は、if文の中でのみ有効
	x := 0
	if x := 2; true {
		fmt.Println(x) // 2
	}
	fmt.Println(x) // 0
}
