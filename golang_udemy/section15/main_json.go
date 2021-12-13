package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// json

type A struct {
}

type User struct {
	ID      int64     `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A,omitempty"`
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@example.com"
	u.Created = time.Now()

	// Marshal 構造体をjsonに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// 構造体に戻す
	fmt.Printf("%T\n", bs)

	u2 := new(User)

	// Unmardhal jsonを構造体に変換
	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)
}
