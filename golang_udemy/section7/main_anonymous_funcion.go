package main

import (
	"fmt"
)

// 関数
// 無名関数
// 一度しか呼ばれない関数や、簡易的な関数の時に使用する

func main() {
	// funcの後に関数名を指定しない
	f := func(x, y int) int {
		return x + y
	}

	i := f(1, 2)
	fmt.Println(i)

	// そのまま実行
	i2 := func(x, y int) int {
		return x + y
	}(1, 3)
	fmt.Println(i2)
}
