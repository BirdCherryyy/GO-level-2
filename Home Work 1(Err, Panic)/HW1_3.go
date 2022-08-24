package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	go func() {
		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
}
