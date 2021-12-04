package main

import "fmt"

// interface
// 異なる型に共通の性質を付ける

type Stringfy interface {
	ToString() string
}

type Pserson struct {
	Name string
	Age  int
}

func (p *Pserson) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Pserson{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-123"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
