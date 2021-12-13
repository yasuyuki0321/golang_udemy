package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `josn:"id"`
	Name    string    `josn:"name"`
	Email   string    `josn:"email"`
	Created time.Time `josn:"created"`
	A       *A        `josn:"A,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		ID   int
		Name string
	}{
		ID:   u.ID,
		Name: "Mr " + u.Name,
	})
	return v, err

}

func (u *User) UnmarshalJSON(b []byte) error {
	type user2 struct {
		Name string
	}
	var u2 user2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		log.Print(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@example.com"
	u.Created = time.Now()

	bs, err := u.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs)
	u2 := new(User)
	if err := json.Unmarshal(bs, &u2); err != nil {
		log.Println(err)
	}
	fmt.Println(u2)
}
