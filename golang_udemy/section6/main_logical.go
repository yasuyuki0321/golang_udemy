package main

import (
	"fmt"
)

// 論理演算子

func main() {
	// and
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	// or
	fmt.Println(true || false == true)
	fmt.Println(false || false == true)

	// not
	fmt.Println(!true)
	fmt.Println(!false)
}
