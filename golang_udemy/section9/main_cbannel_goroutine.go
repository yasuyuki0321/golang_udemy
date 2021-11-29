package main

import (
	"fmt"
	"time"
)

// chanel
// ゴルーチン

func reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 空のチャネルからデータを取ろうとするとdeadlockになる
	// fmt.Println(<-ch1)

	go reciever(ch1)
	go reciever(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
}
