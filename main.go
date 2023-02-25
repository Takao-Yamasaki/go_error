package main

import (
	"errors"
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

// errors.Asを使用する場合
func main() {
	var err1, err2 error

	_, err1 = strconv.Atoi("a") // err1: strconv.NumError型
	err2 = getMyError(err1)     // err2: 内部にstrconv.NumError型を持ったMyError型

	var targetNumErr *strconv.NumError
	// NumError型のエラーインターフェースをNumErrorに変換 -> true
	fmt.Println(errors.As(err1, &targetNumErr))
	// MyError(内部にNumError型)のエラーインターフェースをNumErrorに変換 -> true
	fmt.Println(errors.As(err2, &targetNumErr))

	var targetMyErr *MyError
	// NumError型のエラーインターフェースをMyErrorに変換 -> false
	fmt.Println(errors.As(err1, &targetMyErr))
	// MyError(内部にNumError型)のエラーインターフェースをMyErrorに変換 -> true
	fmt.Println(errors.As(err2, &targetMyErr))
}
