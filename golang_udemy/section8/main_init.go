package main

import "fmt"

// init

// initはmain関数より先に呼ばれる
// main処理より前に初期化処理を行いたい時に使う
func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}
