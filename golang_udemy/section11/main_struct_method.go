package main

import "fmt"

// 構造体
// メソッド

type User struct {
	Name string
	Age  int
}

// レシーバー
// (u User)は型を指定
func (u *User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

// ポインタ型にすることで参照渡しにできる
// 構造体に定義するレシーバはポインタ型にする
func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	// 変数渡し
	user1.SetName("A")
	user1.SayName()

	// 参照渡し
	user1.SetName2("A")
	user1.SayName()

	user2 := &User{Name: "user2"}
	user2.SetName2("B")
	user2.SayName()

}
