package main

import "fmt"

// interface
// カスタムエラー

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaisError() error {
	return &MyError{Message: "カスタムエラーが発生", ErrCode: 1234}
}

func main() {
	err := RaisError()
	fmt.Println(err.Error())

	fmt.Println((err.Error()))

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

}
