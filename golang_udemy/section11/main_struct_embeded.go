package main

import "fmt"

// 構造体
// 埋め込み

// フィールド名と型が同じであることはよくあること
type T struct {
	//	User User
	User // 型名を省略することも可能
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "a"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)

	// 型名を省略したときだけ直接アクセスできる
	fmt.Println(t.Name)
	t.SetName()
	fmt.Println(t.Name)
}
