package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	source := make(chan int)
	double := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(source)
		for _, n := range numbers {
			source <- n
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(double)
		for n := range source {
			double <- n * 2
		}
	}()

	go func() {
		wg.Wait()
		close(double)
	}()

	for result := range double {
		fmt.Println(result)
	}
}
