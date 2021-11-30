package main

import "fmt"

//　ポインタ

func Double(i int) {
	i = i * 2
}

// ポインタを引数として取る
// 参照渡し: スライス、マップ、チャネル
func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)

	// メモリのアドレスを表示
	fmt.Println(&n)

	Double(n)

	// 値渡しなので変数は変わらない
	// Doubleに渡す時にnの値をコピーして渡しているので、別物として扱われる
	fmt.Println(n)

	// ポインタ型を宣言する。
	//  * + 型でポインタ型を宣言
	// ポインタ型の変数に&nでアドレスを入れる
	// pとnは同じアドレスを見ている
	var p *int = &n
	fmt.Println(p)
	// アドレスの値を* + 変数で表示できる
	fmt.Println(*p)

	*p = 300
	fmt.Println(n)

	n = 200
	fmt.Println(*p)

	// Douvlev2にアドレスを渡す
	Doublev2(&n)
	fmt.Println(n)

	Doublev2(p)
	fmt.Println(*p)

	// sliceは参照型なので関数内で書き換えが可能
	sl := []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl)
}
