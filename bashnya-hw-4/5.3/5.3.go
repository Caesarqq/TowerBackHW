package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()

			mu.Lock()
			m[x] = x * x
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	mu.Lock()
	defer mu.Unlock()
	for k, v := range m {
		fmt.Printf("%d -> %d\n", k, v)
	}
}
