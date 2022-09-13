package main

import (
	"fmt"
	"time"
)

type ErrorTime struct { //структура для моей ошибки
	textErr string
	Time    int
}

func New(textErrr string) error {
	return &ErrorTime{
		textErr: "My error",
		Time:    time.Now().Hour(), //фиксация времени
	}
}

func (e *ErrorTime) Error() string {
	return fmt.Sprintf("error: %s\ntime: %v", e.textErr, e.Time)
}

func main() {
	defer func() { // отложеная функция
		if ok := recover(); ok != nil {
			var err error
			err = New("My err")
			fmt.Println(err)
		}
	}()
	CallErr()
}

func CallErr() { //неявная ошибка
	var a int
	fmt.Println(1 / a)
}
