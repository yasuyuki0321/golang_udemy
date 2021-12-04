package main

import "fmt"

// interface
// fmt.stringer

// type Stginger interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v. %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "abc"}
	fmt.Println(p)
}
