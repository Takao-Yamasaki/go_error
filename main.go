package main

import (
	"fmt"
	"strconv"
)

// 内部にエラー入れ子にできるMyError型を用意
type MyError struct {
	InternalError error
}

func (e *MyError) Error() string {
	return e.InternalError.Error()
}

func (e *MyError) Unwrap() error {
	return e.InternalError
}

func getMyError(err error) error {
	return &MyError{err}
}

// 普通の型アサーションを使用する場合
func main() {
	// err1: strconv.NumError型
	// err2: 内部にstrconv.NumError型を持ったMyError型

	var err1, err2 error

	_, err1 = strconv.Atoi("a")
	err2 = getMyError(err1)

	var ok bool

	_, ok = err1.(*strconv.NumError)
	fmt.Println(ok) // true

	// err2をNumErrorにアサーションすることができない
	_, ok = err2.(*strconv.NumError)
	fmt.Println(ok) // false

	_, ok = err1.(*MyError)
	fmt.Println(ok) // false

	_, ok = err2.(*MyError)
	fmt.Println(ok) // true
}
