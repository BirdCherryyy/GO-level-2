package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	f, err := os.Create("trace3.out") //создал файл для отслеживания трассировки
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

	const count = 1000

	var (
		counter int
		wg      sync.WaitGroup
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
