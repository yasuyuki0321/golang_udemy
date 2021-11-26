package main

import(
	"fmt"
)

// 型
// byte(unit8)型

func main() {
	// byte型の配列
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// byteを文字列型に変換
	fmt.Println(string(byteA))

	// 文字列型をbyte型に変換
	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

}
