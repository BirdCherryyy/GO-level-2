package main

import (
	"os"
	"strconv"
)

func main() {
	fpath := "./mil_folder/"
	txt := ".txt"
	var path string
	for i := 0; i < 1000000; i++ {
		path = fpath + strconv.Itoa(i) + txt
		f, err := os.Create(path)
		defer f.Close()
		if err != nil {
			panic(err)
		}
	}
}
