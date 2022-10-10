package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	f, err := os.Create("trace1.out") //создал файл для отслеживания трассировки
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// Your program here
	wg := sync.WaitGroup{}
	wg.Add(2)
	lock := sync.Mutex{}
	go func() {
		defer wg.Done()
		defer lock.Unlock() //Использую мьютекс для того чтобы четко выполнилась задача первого потока
		lock.Lock()
		for i := 0; i < 500; i++ {
			fmt.Println("stream1 i = ", i)
		}
	}()
	go func() {
		defer wg.Done()
		defer lock.Unlock()
		lock.Lock() //Использую мьютекс для того чтобы четко выполнилась задача второго потока
		for i := 0; i < 500; i++ {
			fmt.Println("stream2 i = ", i)
		}
	}()
	wg.Wait()

}
