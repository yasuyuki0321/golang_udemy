package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
		// 簡易的なマッチング
		match, _ := regexp.MatchString("A", "ABC")
		fmt.Println(match)

		// Compile
		re1, _ := regexp.Compile("A")
		match = re1.MatchString("ABC")
		fmt.Println(match)

		// MustCompile
		// コンパイルエラーの時、ランタイムエラーになる
		re2 := regexp.MustCompile("A")
		match = re2.MatchString("ABC")
		fmt.Println(match)
	*/

	/*
		// i: 大文字小文字の区別をしない
		// i-ms: iを有効化して、msは無効化
		re3 := regexp.MustCompile(`(?i)abc`)
		match := re3.MatchString("ABC")
		fmt.Println(match)

		// 完全一致
		re4 := regexp.MustCompile(`^ABC$`)
		match = re4.MatchString("ABC")
		fmt.Println(match)

		match = re4.MatchString("  ABC  ")
		fmt.Println(match)
	*/

	/*
		re6 := regexp.MustCompile("a+b*")
		fmt.Println(re6.MatchString("ab"))
		fmt.Println(re6.MatchString("a"))
		fmt.Println(re6.MatchString("aaaabbbbb"))
		fmt.Println(re6.MatchString("b"))
	*/

	/*
		re8 := regexp.MustCompile(`[XYZ]`)
		fmt.Println(re8.MatchString("Y"))

		re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
		fmt.Println(re9.MatchString("AA_"))
		fmt.Println(re9.MatchString("abcd"))

		re10 := regexp.MustCompile(`[^0-9A-Za-z]`)
		fmt.Println(re10.MatchString("A"))
		fmt.Println(re10.MatchString("あ"))
	*/

	/*
		re11 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
		fmt.Println(re11.MatchString("abcxyz"))
		fmt.Println(re11.MatchString("ABCXYZ"))
		fmt.Println(re11.MatchString("abcXYZ"))
		fmt.Println(re11.MatchString("ABCxyz"))
		fmt.Println(re11.MatchString("ABCabc"))
		fmt.Println(re11.MatchString("ABC"))
	*/

	/*
		re := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
		fmt.Println(re.FindString("AAAABCXYZZZ"))
		// -1 はマッチしたもの全てを返す
		fmt.Println(re.FindAllString("ABCXYZABCXYZ", -1))
	*/

	// re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	// s := `
	// 	0123-456-7889
	// 	111-222-333
	// 	556-787-899
	// `
	// ms := re1.FindAllStringSubmatch(s, -1)
	// fmt.Println(ms)

	// for _, v := range ms {
	// 	fmt.Println(v)
	// }

	// 文字列の置換
	re1 := regexp.MustCompile(`\s+`)
	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ","))

	re2 := regexp.MustCompile(`、|。`)
	fmt.Println(re2.ReplaceAllString("私は、Golangを使用する、プログラマー", ""))

	re3 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re3.Split("ASDFGHabcxyzAWSDGabcXYZGTR", -1))

	re4 := regexp.MustCompile(`\s+`)
	fmt.Println(re4.Split("aaaa  bbb    ccc", -1))

}
