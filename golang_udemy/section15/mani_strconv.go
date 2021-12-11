package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// boolをstringに変換する
	// bt := true
	// fmt.Printf("%T\n", strconv.FormatBool(bt))

	// intをstringに変換する
	// i := strconv.FormatInt(-100, 10)
	// fmt.Printf("%v %T\n", i, i)

	// i2 := strconv.Itoa(100)
	// fmt.Printf("%v %T\n", i2, i2)

	// 浮動小数点を文字列に変換する

	// 文字列を真偽値に変換する
	// bt1, _ := strconv.ParseBool("true")
	// fmt.Printf("%T\n", bt1)
	// bt2, _ := strconv.ParseBool("1")
	// fmt.Printf("%T\n", bt2)
	// bt3, _ := strconv.ParseBool("t")
	// fmt.Printf("%T\n", bt3)
	// bt4, _ := strconv.ParseBool("T")
	// fmt.Printf("%T\n", bt4)
	// bt5, _ := strconv.ParseBool("True")
	// fmt.Printf("%T\n", bt5)
	// bt6, _ := strconv.ParseBool("TRUE")
	// fmt.Printf("%T\n", bt6)
	// fmt.Println(bt1, bt2, bt3, bt4, bt5, bt6)

	// bt1, err := strconv.ParseBool("false")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%T\n", bt1)
	// bt2, _ := strconv.ParseBool("0")
	// fmt.Printf("%T\n", bt2)
	// bt3, _ := strconv.ParseBool("f")
	// fmt.Printf("%T\n", bt3)
	// bt4, _ := strconv.ParseBool("f")
	// fmt.Printf("%T\n", bt4)
	// bt5, _ := strconv.ParseBool("False")
	// fmt.Printf("%T\n", bt5)
	// bt6, _ := strconv.ParseBool("FALSE")
	// fmt.Printf("%T\n", bt6)
	// fmt.Println(bt1, bt2, bt3, bt4, bt5, bt6)

	// 文字列を整数に変換
	// // i3, _ := strconv.ParseInt("1234", 10, 0)
	// // fmt.Printf("%v %T\n", i3, i3)
	// // i4, _ := strconv.ParseInt("-10", 10, 0)
	// // fmt.Printf("%v %T\n", i4, i4)

	// i5, _ := strconv.Atoi("1234")
	// fmt.Println(i5)

	// 文字列を浮動小数点に変換
	fl1, _ := strconv.ParseFloat("3.24", 64)
	fmt.Printf("%v %T\n", fl1, fl1)
	fl2, _ := strconv.ParseFloat(".2", 64)
	fmt.Printf("%v %T\n", fl2, fl2)
	fl3, _ := strconv.ParseFloat("-2", 64)
	fmt.Printf("%v %T\n", fl3, fl3)
	fl4, _ := strconv.ParseFloat("1.2345e8", 64)
	fmt.Printf("%v %T\n", fl4, fl4)
	fl5, _ := strconv.ParseFloat("1.2345E8", 64)
	fmt.Printf("%v %T\n", fl5, fl5)

}
