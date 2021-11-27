package main

import "fmt"

// for
// 繰り返し

func main() {
	/*
		i := 0
		for {
			i++
			if i == 3 {
				break
			}
			fmt.Println("Loop")
		}
	*/

	/*
		// 条件付きfor
		point := 0
		for point < 10 {
			fmt.Println(point)
			point++
		}
	*/

	/*
		//　古典的for文
		for i := 0; i < 10; i++ {
			if i == 3 {
				continue
			}

			if i == 6 {
				break
			}

			fmt.Println(i)
		}
	*/

	/*
		// for文で配列の中身を取り出する
		arr := [3]int{1, 2, 3}

		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	*/

	/*
		// 範囲式for
		// indexと要素を取得できる
		// indexが必要ない場合には _ を使用する
		arr := [3]int{1, 2, 3}

		for i, v := range arr {
			fmt.Println(i, v)
		}
	*/

	/*
		// スライス
		sl := []string{"Python", "PHP", "Go"}
		for i, v := range sl {
			fmt.Println(i, v)
		}
	*/

	// マップ
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// keyだけ
	for k := range m {
		fmt.Println(k)
	}

	// valueだけ
	for _, v := range m {
		fmt.Println(v)
	}
}
