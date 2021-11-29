package main

import "fmt"

// スライス
// for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	// 値だけを使う場合
	for _, v := range sl {
		fmt.Println(v)
	}

	// インデックス番号だけを使う場合
	for i := range sl {
		fmt.Println(i, sl[i])
	}

	// 古典的for
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}
