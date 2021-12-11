package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		fmt.Println(math.Pi)
		fmt.Println(math.Sqrt(2))

		// float32で表現可能な最大値
		fmt.Println(math.MaxFloat32)
		// float32で表現可能な0でない最小値
		fmt.Println(math.SmallestNonzeroFloat32)

		// float64で表現可能な最大値
		fmt.Println(math.MaxFloat64)
		// float64で表現可能な0でない最小値
		fmt.Println(math.SmallestNonzeroFloat64)

		fmt.Println(math.MaxInt64)
		fmt.Println(math.MinInt64)
	*/

	/*
		// 絶対値
		fmt.Println(math.Abs(100))
		fmt.Println(math.Abs(-100))

		// 累乗
		fmt.Println(math.Pow(0, 2))
		fmt.Println(math.Pow(2, 2))

		// 平方根、立方根
		fmt.Println(math.Sqrt(2))
		fmt.Println(math.Cbrt(8))

		// 最大値、最小値
		fmt.Println("Max:", math.Max(10, 2))
		fmt.Println("Min:", math.Min(10, 2))
	*/

	// 小数点
	// 単純な切り捨て
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	// 引数の数値より小さい最大の整数
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	// 引数の数値より大きい最大の整数
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))
}
