package main

import (
	"fmt"
	"time"
)

// channel
// close

func reciever(name string, c chan int) {
	for {
		i, ok := <-c
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	// チャネルは作成時はopenの状態
	ch1 := make(chan int, 2)

	/*
		ch1 <- 1
		close(ch1)

		// closeしたchannelにデータを送ることはできない
		// send on closed channel
		// ch1 <- 1

		// closeされたchannelからデータを取得することは可能
		//fmt.Println(<-ch1)

		// 1つめの変数は受け取った値
		// 2つめの変数はチャネルのopen状態を表す
		// バッファ内にデータが空で、channelがcloseされた場合にflaseとなる
		i, ok := <-ch1
		fmt.Println(i, ok)
		i2, ok := <-ch1
		fmt.Println(i2, ok)
	*/

	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}
