package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		fmt.Println("表示")

		// fmt標準
		fmt.Print("Hello\n")
		// 開業
		fmt.Println(("Hello!"))
	*/

	/*
		fmt.Println("Hello", "world!")
		fmt.Println("Hello", "world!")
	*/

	/*
		// フォーマットを指定
		fmt.Printf("%s\n", "Hello")
		fmt.Printf("%#v\n", "Hello")
	*/

	/*
		// Sprint は、出漁ではなくフォーマットした結果を文字列で返す
		s := fmt.Sprint("Hello")
		s1 := fmt.Sprintf("%v\n", "Hello")
		s2 := fmt.Sprintln("Hello")

		fmt.Println(s)
		fmt.Println(s1)
		fmt.Println(s2)
	*/

	// Fprint 書き込み先を指定
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test1.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint")

}
