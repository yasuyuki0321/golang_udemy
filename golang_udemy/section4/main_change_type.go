package main

import (
	"fmt"
	"strconv"
)

// 型変換

func main() {
	/*
		var i int = 1
		// int → float64への変換
		fl64 := float64(i)

		fmt.Println((fl64))
		fmt.Printf("i = %T\n", i)
		fmt.Printf("fl64 = %T\n", fl64)

		i2 := int(fl64)
		fmt.Printf("i2 = %T\n", i2)

		// float64をint型に変換すると小数点以下は切り捨てられる
		fl := 10.5
		i3 := int(fl)
		fmt.Println(i3)
		fmt.Printf("i3 = %T\n", i3)
	*/

	// 文字列型を数値型に変換する
	// strconvパッケージを使用する必要がある
	var s string = "100"
	fmt.Printf("s = %T\n", s)

	// 2つめの戻り値を _ にすることで戻り値を破棄できる
	// goでは宣言した変数は必ず指定する必要があるので、使用しない場合には破棄する
	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	// 2つめの戻り値の err を処理する場合
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	// int型を文字列型に変換
	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	// 文字列をbyte配列に変換する
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)
	fmt.Printf("b = %T\n", b)

	// byte配列を文字列型に変換する
	h2 := string(b)
	fmt.Println(h2)
	fmt.Printf("h2 = %T\n", h2)
}
