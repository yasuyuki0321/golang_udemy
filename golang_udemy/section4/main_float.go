package main

import (
	"fmt"
)

// 型
// 浮動小数点

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Printf("%T %T\n", fl64, fl)
	fmt.Println(fl64 + fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Printf("%T\n", float64(fl32))

	fmt.Printf("%T\n", float32(fl64))

	zero := 0.0
	pinf := 1.0 / zero
	// 正の無限大
	fmt.Println(pinf)

	ninf := -1.0 / zero
	// 負の無限大
	fmt.Println(ninf)

	nan := zero / zero
	// 非数
	fmt.Println((nan))
}
