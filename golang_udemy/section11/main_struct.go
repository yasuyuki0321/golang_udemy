package main

import "fmt"

// 構造体

// type + 構造体名 + struct で構造体の定義が可能
type User struct {
	// フィールド
	Name string
	Age  int
	// まとめてていぎすることも可能
	// x, y int
}

// userという変数でUser型を渡す
// 構造体は値渡し
func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

// 参照渡し
func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	// 明示的な宣言
	// user1はUser型
	var user1 User
	// 構造体は{}で表示される
	// { 0} 空文字(stringの初期値)、0(intの初期値が表示される)
	fmt.Println(user1)

	// 各フィールドに値を設定
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	// 暗黙的な宣言
	user2 := User{}
	fmt.Println(user2)

	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	// フィールドを指定しなくても代入可能
	// 構造体で定義したフィールドの順番で指定する必要がある
	user4 := User{"user4", 40}
	fmt.Println(user4)

	// 型エラーになる
	// user5 := User{30, "user5"}
	// fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	// newを使うと構造体のポインタ型が返される
	// &{ 0}
	user7 := new(User)
	fmt.Println(user7)
	fmt.Println(*user7)

	// アドレス演算子を付けるとポインタ型で宣言できる
	// newと同じ
	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)

}
