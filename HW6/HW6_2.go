package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	f, err := os.Create("trace2.out") //создал файл для отслеживания трассировки
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

	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(0)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		log.Println("I'm working!")
	}()

	for i := 0; i < 5000000000; i++ {
		if i%1000000 == 0 {
			runtime.Gosched() //явный вызов планировщика каждые 1000000 операций
			fmt.Println(":)")
		}
	}
	wg.Wait()
}
