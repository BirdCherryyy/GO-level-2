package main

import (
	"fmt"
	"sync"
)

func main() {
	n_stream(1000)
	nStreamWithLock(1000)
}

func n_stream(n int) { // запуск n потоков и ожидание их завершения
	wg := sync.WaitGroup{}
	mystream := n
	wg.Add(mystream)
	for i := 0; i < n; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("n_stream is done")
}

func nStreamWithLock(n int) { // запуск n потоков и ожидание их завершения
	var (
		wg       = sync.WaitGroup{}
		mutex    sync.Mutex
		mystream = n
	)
	wg.Add(mystream)
	for i := 0; i < n; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("n_stream is done")
}
