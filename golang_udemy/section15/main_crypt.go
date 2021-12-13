package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	// 任意の文字列からmd5ハッシュを生成する
	h := md5.New()

	io.WriteString(h, "ABCD")
	fmt.Println(h.Sum(nil))

	// 16進数の文字列を生成
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)
}
