package main

import "fmt"

func main() {
	m := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// keyを使用しない場合
	for _, v := range m {
		fmt.Println(v)
	}

	// keyだけ取り出す
	for k := range m {
		fmt.Println(k)
	}
}
