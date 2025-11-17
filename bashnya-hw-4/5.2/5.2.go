package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <количество_воркеров>")
		os.Exit(1)
	}
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Ошибка: укажите положительное число воркеров")
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dataChan := make(chan int)
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dataChan <- i
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println(id)
					return
				case data := <-dataChan:
					fmt.Println(id, data)
				}
			}
		}(i)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	cancel()
	close(dataChan)
	wg.Wait()
	fmt.Println("Всё ок")
}
