package main

import (
	"fmt"
	"reflect"
)

type in struct {
	key int
}

type values map[string]interface{}

func main() {
	var (
		a in
		c values
	)
	c["one"] = 1
	enterValue(a, c)
	fmt.Println(a.key)
}

func enterValue(input in, value values) error {
	var ok error
	rInput := reflect.ValueOf(input.key).Elem()
	switch value["one"].(type) {
	case int:
		rInput.SetInt(1)
	case values:
		rInput.SetInt(2)
	default:
		fmt.Println("s: ")
	}

	return ok
}
