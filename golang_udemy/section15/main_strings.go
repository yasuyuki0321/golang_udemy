package main

import (
	"fmt"
	"strings"
)

func main() {

	// 文字列を結合する
	// s1 := strings.Join([]string{"A", "B", "C"}, ",")
	// s2 := strings.Join([]string{"A", "B", "C"}, "")

	// fmt.Println(s1, s2)

	// 文字列の検索
	// 該当の文字列がない場合には-1が返る
	i1 := strings.Index("ABCED", "A")
	i2 := strings.Index("ABCEDE", "F")
	fmt.Println(i1, i2)

	i3 := strings.Index("ABCEDABCDE", "BC")
	i4 := strings.LastIndex("ABCEDABCDE", "BC")
	fmt.Println(i3, i4)

	i5 := strings.IndexAny("ABC", "ABC")
	i6 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i5, i6)

	// 指定した文字列で開始されるかを確認
	b1 := strings.HasPrefix("ABC", "A")
	b2 := strings.HasPrefix("ABC", "C")
	fmt.Println(b1, b2)

	// 指定した文字列が含まれるか
	b3 := strings.Contains("ABC", "B")
	b4 := strings.ContainsAny("ABC", "AD")
	fmt.Println(b3, b4)

	// 指定した文字列がないかい出現するか
	// ""の場合文字列の長さ+1
	i7 := strings.Count("ABCDABCD", "D")
	i8 := strings.Count("ABCDABCD", "")
	fmt.Println(i7, i8)

	// 文字列を繰り返し結合する
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4)

	// 文字列の置換
	// -1は該当する場所すべて
	s5 := strings.Replace("AAAAAA", "A", "B", 2)
	s6 := strings.Replace("AAAAAA", "A", "B", -1)
	fmt.Println(s5, s6)

	// 文字列の分割
	s7 := strings.Split("A,B,C,D,E", ",")
	fmt.Println(s7)
	// カンマを残して分割
	s8 := strings.SplitAfter("A,B,C,D,E", ",")
	fmt.Println(s8)
	s9 := strings.SplitN("A,B,C,D,E", ",", 2)
	fmt.Println(s9)
	s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
	fmt.Println(s10)

	// 大文字小文字の変換
	s11 := strings.ToLower("ABCD")
	s12 := strings.ToLower("E")
	fmt.Println(s11, s12)

	s13 := strings.ToUpper("abcd")
	s14 := strings.ToUpper("e")
	fmt.Println(s13, s14)

	// 文字列から空白を取り除く
	s15 := strings.TrimSpace("   -   ABC   -   ")
	fmt.Println(s15)
	// 全角も対象
	s16 := strings.TrimSpace("　　-  ABC  - 　　")
	fmt.Println(s16)

	// 文字列からスペースで区切られたフィールドを取り出す
	s18 := strings.Fields("a b c")
	fmt.Println(s18)
	fmt.Println(s18[0])

}
