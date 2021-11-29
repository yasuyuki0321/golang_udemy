package main

import "fmt"

// channel
// 複数のコルーチン間でのデータの受け渡しをするために設計されたデータ構造
// 宣言、操作

func main() {
	// int型のデータを保持できるチャネル
	// 受信送信を指定しない場合には、送受信可能なチャネルになる
	var ch1 chan int

	/*
		// 受信専用のチャネル
		var ch2 <-chan int

		// 送信専用のチャネル
		var ch3 chan<- int
	*/

	// チャネルの初期化
	// チャネルの読み込み/書き込みができるようになる
	ch1 = make(chan int)
	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	// 容量が5のチャネルを作成
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	// データの送信
	// ch3 にデータを送信
	ch3 <- 1
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len:", len(ch3))

	// データの受信
	i := <-ch3
	fmt.Println(i)
	// チャネルの要素数を確認すると数が減っている
	fmt.Println("len:", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len:", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("len:", len(ch3))

	// チャネルにはキューの性質がある
	// 先入先出

	// バッファサイズを超えるとdeadlockになる
	// all goroutines are asleep - deadlock!
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6
	fmt.Println(cap(ch3))
}
