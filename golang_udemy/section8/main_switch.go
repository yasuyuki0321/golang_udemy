package main

import "fmt"

// switch
// 式switch

// switchの型とcaseの型は同じでないとエラーになる
func main() {
	/*
		n := 5
		switch n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("I doon't know")
		}
	*/

	/*
		// 変数が使えるのはswitchの中だけ
		switch n := 2; n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("I doon't know")
		}
		// undefined: n
		// fmt.Println(n)
	*/

	/*
		// case部分に演算子を使用することができる
		n := 6
		switch {
		case n > 0 && n < 4:
			fmt.Println("0 < n < 4")
		case n > 3 && n < 7:
			fmt.Println("3 < n < 7")
		}
	*/

	// 列挙型の式と演算子を使用した式の混在はできない
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	// case n > 3 && n < 7:
	// 	fmt.Println("3 < n < 7")
	// }
	default:
		fmt.Println("I doon't know")
	}
}
