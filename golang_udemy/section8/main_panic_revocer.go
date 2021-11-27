package main

import "fmt"

// panic recover
// panicで強制終了させるよりより、エラーハンドリングちゃんとすることが推奨
func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime errot")
	fmt.Println("START")
}
