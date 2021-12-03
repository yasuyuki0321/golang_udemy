package main

import "fmt"

// 構造体
// map

type User struct {
	Name string
	Age  int
}

func main() {
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}

	fmt.Println(m)

	m2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}
	fmt.Println(m2)

	m3 := make(map[int]User, 0)
	fmt.Println(m3)
	m3[1] = User{"user3", 30}
	fmt.Println(m3)

	for _, v := range m {
		fmt.Println(v)
	}

}
