package main

import "fmt"

// マップ

func main() {
	// 明示的宣言
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	// 暗黙的宣言
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	// 開業してもOK
	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	// make関数
	m4 := make(map[int]string)
	fmt.Println(m4)

	// 値の設定
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	// 値の取り出し
	fmt.Println(m["A"])
	fmt.Println(m4[2])

	// 登録されていない場合には初期値が取得される
	fmt.Println(m4[3])

	// 2つめの戻り値に値が取得できたかが設定される
	s, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok)

	m4[2] = "US"
	fmt.Println(m4[2])
	m4[3] = "CHINA"
	fmt.Println(m4)

	// 削除
	// delete(マップ, キー)
	delete(m4, 3)
	fmt.Println(m4)

	// 要素数を数える
	fmt.Println(len(m4))

}
