package main

import "fmt"

// channel
// for

func main() {
	ch1 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	// チャネルでforを使う場合はcloseしてからforでデータを取得する
	close(ch1)

	for i := range ch1 {
		fmt.Println(i)
	}
}