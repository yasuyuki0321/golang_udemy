package main

import "fmt"

// switch
// 型スイッチ

// interface型の変数を使用すると様々な型の変数を取ることができる
// 渡された引数はinterface型になるので、計算等はできない
func anything(a interface{}) {
	// fmt.Println(a)

	// 受け取った引数の方に応じて処理を行う
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	// 文字列型
	anything("aaa")
	// 数値型
	anything(1)

	var x interface{} = 3
	// x(interface型)をint型で復元 型アサーション
	// i := x.(int)
	// fmt.Println(i + 2)
	// xはinterface型になるのでエラーになる
	// invalid operation: x + 2 (mismatched types interface {} and int)
	// fmt.Println(x + 2)

	// xはint型の変数なので、float型で復元できないためpanicになる
	// f := x.(float64)
	// fmt.Println(f)

	// 2つの返り値は変換に失敗した時にfalseが返ってくる
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	// 型変換に成功したので2つめの引数はtrueになる
	i2, isInt := x.(int)
	fmt.Println(i2, isInt)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isStging := x.(string); isStging {
		fmt.Println(s, isStging)
	} else {
		fmt.Println("I don't know")
	}

	// 型スイッチ
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	// 値を使用する場合
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
}
