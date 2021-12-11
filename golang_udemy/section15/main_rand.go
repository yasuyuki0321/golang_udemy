package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(256)

	// seedを設定しているので、毎回同じ乱数が生成される
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	// 現在時刻をシードに使った疑似乱数の生成
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	// 0から99の擬似乱数の生成
	fmt.Println(rand.Intn(100))
	// int型の疑似乱数
	fmt.Println(rand.Int())
	fmt.Println(rand.Int31())
	fmt.Println(rand.Int63())
	fmt.Println(rand.Uint32())
}
