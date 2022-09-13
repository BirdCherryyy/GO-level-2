package main

import (
	"fmt"
	"time"
)

func main() {
	////////////////////////////////////////////////////////////////////////////////////////////
	fmt.Println("Задание №1:")
	var result = int(0)
	var workers = make(chan struct{}, 1000)
	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}
		go func(job int) {
			defer func() {
				<-workers
			}()
			fmt.Println(job)
			result++
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Result = %d,\n\n", result) //вывод конечного результата инкрементов воркеров
	////////////////////////////////////////////////////////////////////////////////////////////
	fmt.Println("Задание №2:")
	ch2 := make(chan string)

	go func() {
		for i := 0; i != 50; i++ {
			ch2 <- "NOT SIGTERM1"
		}
	}()
	go func() {
		ch2 <- "SIGTERM"
	}()
	////////////////////////////////////////////////////////////////////////////////////////////
	for {
		select {
		case val := <-ch2:
			{
				if val == "SIGTERM" {
					fmt.Println("correct return")
					return
				} else {
					fmt.Println(val)
				}
			}
		case <-time.After(1 * time.Second):
			fmt.Println("TimeOut")
		}
		fmt.Println("End")
	}

}
