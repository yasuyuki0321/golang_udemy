package main

import (
	"fmt"
	"time"
)

// 並列処理
// go goroutin

// go文を使うことで簡単に並列処理が実行できる
func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub()
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
