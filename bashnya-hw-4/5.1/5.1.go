package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	resultChan := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultChan <- n * n
		}(num)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	sum := 0
	for sq := range resultChan {
		sum += sq
	}

	fmt.Print(sum)
}
